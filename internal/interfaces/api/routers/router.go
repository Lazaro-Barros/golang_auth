package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/severusTI/auth_golang/internal/interfaces/api/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupUserRoutes(userHandlers handlers.UserHandlers) *gin.Engine {

	r := gin.Default()

	baseGroup := r.Group("/api")

	usersRouter := baseGroup.Group("users")

	// Rota para criar um usuário
	usersRouter.POST("", userHandlers.CreateUser)

	// Rota para obter um usuário por ID
	usersRouter.GET("/:id", userHandlers.GetUser)

	// Rota para obter lista de usuários
	usersRouter.GET("/list", userHandlers.LisUsers)

	// Rota para atualizar um usuário
	usersRouter.PUT("/:id", userHandlers.UpdateUser)

	// Rota para excluir um usuário
	usersRouter.DELETE("/:id", userHandlers.DeleteUser)

	// add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
