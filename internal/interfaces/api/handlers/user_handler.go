package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/severusTI/auth_golang/internal/usecases"
	"github.com/severusTI/auth_golang/internal/usecases/dtos"
	"github.com/severusTI/auth_golang/pkg/ops"
)

type UserHandlers struct {
	userUseCases usecases.UserUC
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

	if err = uh.userUseCases.CreateUser(&user); err != nil {
		ops.Handling(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// GetUser godoc
// @Summary get a user
// @Description Get a user by id from the database
// @Param userID path string true "Get user"
// @Produce application/json
// @Tags user
// @Success 200 {object} dtos.ResUser{}
// @Router /users/{userID} [GET]
func (uh *UserHandlers) GetUser(c *gin.Context) {
	var (
		err    error
		userID = c.Param("id")
	)
	user, err := uh.userUseCases.GetUser(&userID)
	if err != nil {
		ops.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, *user)
}

// ListUsers godoc
// @Summary get all users
// @Description get all users from the database
// @Produce application/json
// @Tags user
// @Success 200 {object} []dtos.ResUser{}
// @Router /users/list [GET]
func (uh *UserHandlers) LisUsers(c *gin.Context) {
	var (
		err error
	)
	users, err := uh.userUseCases.ListUsers()
	if err != nil {
		ops.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
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

	if err = uh.userUseCases.UpdateUser(userID, &user); err != nil {
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

	if err = uh.userUseCases.DeleteUser(&userID); err != nil {
		ops.Handling(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewUserHandlers(userUseCases usecases.UserUC) (obj *UserHandlers) {
	obj = &UserHandlers{
		userUseCases: userUseCases,
	}
	return
}
