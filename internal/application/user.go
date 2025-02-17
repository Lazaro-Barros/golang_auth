package application

import (
	"github.com/severusTI/auth_golang/internal/application/dtos"
	domain "github.com/severusTI/auth_golang/internal/domain/entities"
	"github.com/severusTI/auth_golang/internal/interfaces/persistance/repositories"
	"github.com/severusTI/auth_golang/pkg/ops"
)

type IUserApplication interface {
	CreateUser(reqUser *dtos.ReqUser) (err error)
	UpdateUser(userID string, reqUser *dtos.ReqUser) (err error)
	DeleteUser(userID *string) (err error)
}

type UserApplication struct {
	userRepository repositories.IUserRepository
}

func NewUserApplication(userRepository repositories.IUserRepository) IUserApplication {
	return &UserApplication{
		userRepository: userRepository,
	}
}

func (obj *UserApplication) CreateUser(reqUser *dtos.ReqUser) (err error) {
	user, err := domain.NewUser(reqUser.Email, reqUser.InputedPassword, reqUser.Name, reqUser.PhoneNumber)
	if err != nil {
		return
	}

	if err := obj.userRepository.CreateUser(user); err != nil {
		return ops.Err(err)
	}
	return
}

func (obj *UserApplication) UpdateUser(userID string, reqUser *dtos.ReqUser) (err error) {
	user, err := domain.NewUser(reqUser.Email, reqUser.InputedPassword, reqUser.Name, reqUser.PhoneNumber)
	if err != nil {
		return
	}

	if err := obj.userRepository.UpdateUser(&userID, user); err != nil {
		return ops.Err(err)
	}

	return

}
func (obj *UserApplication) DeleteUser(userID *string) (err error) {
	if err := obj.userRepository.DeleteUser(userID); err != nil {
		return ops.Err(err)
	}
	return
}
