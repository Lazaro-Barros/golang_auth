package application

import (
	"github.com/severusTI/auth_golang/internal/application/dtos"
	"github.com/severusTI/auth_golang/internal/domain"
	"github.com/severusTI/auth_golang/internal/interfaces/persistance/repositories"
	"github.com/severusTI/auth_golang/pkg/ops"
)

type IUserApplication interface {
	CreateUser(reqUser *dtos.ReqUser) (err error)
	GetUser(userID *string) (resUser *dtos.ResUser, err error)
	ListUsers() (resUsers []dtos.ResUser, err error)
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

func (obj *UserApplication) GetUser(userID *string) (resUser *dtos.ResUser, err error) {
	user, err := obj.userRepository.GetUser(userID)
	if err != nil {
		return resUser, ops.Err(err)
	}
	return &dtos.ResUser{
		ID:          user.Id(),
		Name:        user.Name(),
		PhoneNumber: user.PhoneNumber(),
		Email:       user.Email(),
	}, err
}

func (obj *UserApplication) ListUsers() (resUsers []dtos.ResUser, err error) {
	users, err := obj.userRepository.ListUsers()
	if err != nil {
		return resUsers, ops.Err(err)
	}

	resUsers = []dtos.ResUser{}
	for i := range users {
		resUser := dtos.ResUser{
			ID:          users[i].Id(),
			Name:        users[i].Name(),
			PhoneNumber: users[i].PhoneNumber(),
			Email:       users[i].Email(),
		}
		resUsers = append(resUsers, resUser)
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
