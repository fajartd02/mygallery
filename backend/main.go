package main

import (
	"github.com/fajartd02/mygallery/backend/config"
	"github.com/fajartd02/mygallery/backend/core/entity"
	"github.com/fajartd02/mygallery/backend/core/module"
	"github.com/fajartd02/mygallery/backend/handler"
	memoryrepository "github.com/fajartd02/mygallery/backend/repository/memory"
	userrepository "github.com/fajartd02/mygallery/backend/repository/user"
	"github.com/fajartd02/mygallery/backend/routes"
)

func main() {

	db := config.InitDB()
	// cfg := config.InitConfig()
	db.AutoMigrate(&entity.User{}, &entity.Memory{})

	userRepo := userrepository.New()
	memoryRepo := memoryrepository.New()

	userUc := module.NewUserUsecase(userRepo)
	memoryUc := module.NewMemoryUseCase(memoryRepo)

	userHdl := handler.NewUserHandler(userUc)
	memoryHdl := handler.NewMemoryHandler(memoryUc)

	r := routes.SetupRoutes(db, *userHdl, *memoryHdl)

	r.Run(":8080")
}
