package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/severusTI/auth_golang/internal/application"
	"github.com/severusTI/auth_golang/internal/application/dtos"
)

type LoginHandler struct {
	loginApp application.ILoginApplication
}

func NewLoginHandler(loginApp application.ILoginApplication) (obj *LoginHandler) {
	obj = &LoginHandler{
		loginApp: loginApp,
	}
	return
}

// LoginUser godoc
// @Summary User login
// @Description Authenticates a user and returns a JWT token
// @Param credentials body dtos.ReqLogin true "User credentials"
// @Produce application/json
// @Tags user
// @Success 200 {object} dtos.ResLogin
// @Router /users/login [POST]
func (obj *LoginHandler) LoginUser(c *gin.Context) {
	var loginData dtos.ReqLogin
	var err error

	if err = c.ShouldBindJSON(&loginData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := obj.loginApp.Login(&loginData)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// Retorna o token ao usuário
	c.JSON(http.StatusOK, token)
}
