package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fajartd02/mygallery/core/module"
	"github.com/fajartd02/mygallery/handler"
)

// -- DEV ONLY -- cors setup
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func SetupRoutes(
	db *gorm.DB,
	// cfg config.Config,
	userHdl handler.UserHandler,
	memoryHdl handler.MemoryHandler,
) *gin.Engine {

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		// c.Set("host", cfg.Host)
	})

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "okss")
	})

	// -- DEV ONLY --
	r.Use(CORSMiddleware())

	apiV1 := r.Group("/api/v1")
	{
		// Authentication and Authorization routes
		apiV1.POST("/login", userHdl.Login)
		apiV1.POST("/register", userHdl.Register)

		// User routes
		user := apiV1.Group("/users")
		user.Use(module.IsAuthorized())
		user.GET("", userHdl.GetAll)
		user.GET("/:id", userHdl.GetSingle)
		user.PUT("/:id", userHdl.Update)
		user.DELETE("/:id", userHdl.Delete)

		// File routes
		file := apiV1.Group("/files")
		file.StaticFS("/static", http.Dir("images"))

		memory := apiV1.Group("/memories")
		memory.Use(module.IsAuthorized())
		memory.GET("", memoryHdl.GetAll)
		memory.GET("/:memoryId", memoryHdl.GetById)
		memory.POST("", memoryHdl.CreateMemory)
		memory.PUT("/:memoryId", memoryHdl.UpdateMemory)
		memory.DELETE("/:memoryId", memoryHdl.DeleteMemory)
		memory.POST("/list", memoryHdl.ListMemory)

	}

	return r
}
