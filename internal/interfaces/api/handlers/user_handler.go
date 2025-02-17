package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/severusTI/auth_golang/internal/application"
	"github.com/severusTI/auth_golang/internal/application/dtos"
	"github.com/severusTI/auth_golang/pkg/ops"
)

type UserHandlers struct {
	userApp application.IUserApplication
}

func NewUserHandler(userApp application.IUserApplication) (obj *UserHandlers) {
	obj = &UserHandlers{
		userApp: userApp,
	}
	return
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user in the database
// @Param user body dtos.ReqUser true "Create user"
// @Produce application/json
// @Tags user
// @Success 201
// @Router /users [POST]
func (uh *UserHandlers) CreateUser(c *gin.Context) {
	var (
		user dtos.ReqUser
		err  error
	)

	if err = c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = uh.userApp.CreateUser(&user); err != nil {
		ops.Handling(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user in the database
// @Param userID path string true "Update user"
// @Param user body dtos.ReqUser true "Update user"
// @Produce application/json
// @Tags user
// @Success 204
// @Router /users/{userID} [PUT]
func (uh *UserHandlers) UpdateUser(c *gin.Context) {
	var (
		user   dtos.ReqUser
		userID = c.Param("id")
		err    error
	)

	if err = c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = uh.userApp.UpdateUser(userID, &user); err != nil {
		ops.Handling(c, err)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user in the database
// @Param userID path string true "Delete user"
// @Produce application/json
// @Tags user
// @Success 204
// @Router /users/{userID} [DELETE]
func (uh *UserHandlers) DeleteUser(c *gin.Context) {
	var (
		err    error
		userID = c.Param("id")
	)

	if err = uh.userApp.DeleteUser(&userID); err != nil {
		ops.Handling(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
