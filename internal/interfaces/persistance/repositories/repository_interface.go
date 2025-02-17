package repositories

import "github.com/severusTI/auth_golang/internal/domain/entities"

type IUserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByEmail(email *string) (*entities.User, error)
	UpdateUser(userID *string, user *entities.User) error
	DeleteUser(userID *string) error
}
