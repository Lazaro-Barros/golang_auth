package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/severusTI/auth_golang/internal/interfaces/api/handlers"
)

func SetupUserRoutes(userHandlers *handlers.UserHandlers, loginHandler *handlers.LoginHandler) *gin.Engine {

	r := gin.Default()
	baseGroup := r.Group("/api")
	usersRouter := baseGroup.Group("users")
	// Rota para criar um usuário
	usersRouter.POST("", userHandlers.CreateUser)
	// Rota para atualizar um usuário
	usersRouter.PUT("/:id", userHandlers.UpdateUser)
	// Rota para excluir um usuário
	usersRouter.DELETE("/:id", userHandlers.DeleteUser)

	// Rota para fazer login
	usersRouter.POST("/login", loginHandler.LoginUser)

	// add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
