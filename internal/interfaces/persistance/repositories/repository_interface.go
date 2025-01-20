package repositories

import "github.com/severusTI/auth_golang/internal/domain"

type IUserRepository interface {
	CreateUser(user *domain.User) error
	GetUser(userID *string) (*domain.User, error)
	ListUsers() ([]domain.User, error)
	UpdateUser(userID *string, user *domain.User) error
	DeleteUser(userID *string) error
}
