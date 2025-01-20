package main

import (
	"log"

	_ "github.com/severusTI/auth_golang/docs"
	"github.com/severusTI/auth_golang/internal/application"
	"github.com/severusTI/auth_golang/internal/interfaces/api/handlers"
	"github.com/severusTI/auth_golang/internal/interfaces/api/routers"
	"github.com/severusTI/auth_golang/internal/interfaces/persistance/repositories"
	"github.com/severusTI/auth_golang/pkg/database"
	envload "github.com/severusTI/auth_golang/pkg/env_load"
)

// @title CRUD user API
// @version 1.0
// @description A CRUD user service

// @host localhost:8080
// @BasePath /api
func main() {
	envload.Init()
	db := database.InitDBConnection(envload.Get())
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userusecase := application.NewUserApplication(userRepo)
	userHandlers := handlers.NewUserHandlers(userusecase)

	r := routers.SetupUserRoutes(*userHandlers)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
