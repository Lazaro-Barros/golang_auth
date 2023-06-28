package usecases

import (
	"github.com/severusTI/auth_golang/internal/domain"
	"github.com/severusTI/auth_golang/internal/interfaces/persistance/repositories"
	"github.com/severusTI/auth_golang/internal/usecases/dtos"
	"github.com/severusTI/auth_golang/pkg/ops"
)

type UserUC struct {
	userRepository repositories.IUserRepository
}

func NewUserUC(userRepository repositories.IUserRepository) (obj *UserUC) {
	obj = &UserUC{
		userRepository: userRepository,
	}
	return
}

func (uc *UserUC) CreateUser(reqUser *dtos.ReqUser) (err error) {
	user, err := domain.NewUser(reqUser.Email, reqUser.InputedPassword, reqUser.Name, reqUser.PhoneNumber)
	if err != nil {
		return
	}

	tx, err := uc.userRepository.BeginTransaction()
	if err != nil {
		return ops.Err(err)
	}
	defer tx.Rollback()

	if err := uc.userRepository.CreateUser(user); err != nil {
		return ops.Err(err)
	}

	if tx.Commit(); err != nil {
		return ops.Err(err)
	}

	return
}

func (uc *UserUC) GetUser(userID *string) (resUser *dtos.ResUser, err error) {
	tx, err := uc.userRepository.BeginTransaction()
	if err != nil {
		return resUser, ops.Err(err)
	}
	defer tx.Rollback()

	user, err := uc.userRepository.GetUser(userID)
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

func (uc *UserUC) ListUsers() (resUsers []dtos.ResUser, err error) {
	tx, err := uc.userRepository.BeginTransaction()
	if err != nil {
		return resUsers, ops.Err(err)
	}
	defer tx.Rollback()

	users, err := uc.userRepository.ListUsers()
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
func (uc *UserUC) UpdateUser(userID string, reqUser *dtos.ReqUser) (err error) {
	user, err := domain.NewUser(reqUser.Email, reqUser.InputedPassword, reqUser.Name, reqUser.PhoneNumber)
	if err != nil {
		return
	}

	tx, err := uc.userRepository.BeginTransaction()
	if err != nil {
		return ops.Err(err)
	}
	defer tx.Rollback()

	if err := uc.userRepository.UpdateUser(&userID, user); err != nil {
		return ops.Err(err)
	}

	if tx.Commit(); err != nil {
		return ops.Err(err)
	}

	return

}
func (uc *UserUC) DeleteUser(userID *string) (err error) {
	tx, err := uc.userRepository.BeginTransaction()
	if err != nil {
		return ops.Err(err)
	}
	defer tx.Rollback()

	if err := uc.userRepository.DeleteUser(userID); err != nil {
		return ops.Err(err)
	}

	if tx.Commit(); err != nil {
		return ops.Err(err)
	}

	return
}
