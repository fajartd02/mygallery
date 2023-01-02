package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/fajartd02/mygallery/core/module"
	"github.com/fajartd02/mygallery/handler"
)

func SetupRoutes(
	db *gorm.DB,
	// cfg config.Config,
	userHdl handler.UserHandler,
	memoryHdl handler.MemoryHandler,
	commentHdl handler.CommentHandler,
	fileHdl handler.FileHandler,
	userVoteHdl handler.UserVoteHandler,
	reportHdl handler.ReportHandler) *gin.Engine {

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
		file.POST("", fileHdl.UploadFile)
		file.StaticFS("/static", http.Dir("images"))

		// Comment routes
		comment := apiV1.Group("/comments")
		comment.Use(module.IsAuthorized())
		comment.POST("/:reportId", commentHdl.Create)
		comment.GET("/:reportId", commentHdl.GetAll)

		// Report routes
		report := apiV1.Group("/reports")
		report.Use(module.IsAuthorized())
		report.POST("/:reportId/votes", userVoteHdl.VotingReport)
		report.GET("/:reportId/votes", userVoteHdl.GetVotesInReport)
		report.GET("/:reportId", reportHdl.GetReport)
		report.GET("", reportHdl.GetReports)
		report.POST("", reportHdl.CreateReport)

		memory := apiV1.Group("/memories")
		memory.GET("", memoryHdl.GetAll)
		memory.POST("", memoryHdl.CreateMemory)
		memory.PUT("/:memoryId", memoryHdl.UpdateMemory)
		memory.DELETE("/:memoryId", memoryHdl.DeleteMemory)

	}

	return r
}
