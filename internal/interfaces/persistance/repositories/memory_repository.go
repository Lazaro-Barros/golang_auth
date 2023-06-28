package repositories

import (
	"github.com/severusTI/auth_golang/internal/domain"
	"github.com/severusTI/auth_golang/pkg/ops"
)

type UserRepositoryMemory struct {
	Users []domain.User
}

type MockedTransaction struct {
}

func (MockedTransaction) Commit() error {
	return nil
}
func (MockedTransaction) Rollback() error {
	return nil
}

func NewUserRepositoryMemory() *UserRepositoryMemory {
	repo := UserRepositoryMemory{}
	user1, _ := domain.NewUser("mocked1@user.com", "!@#123Mockedpassword", "mocker user 1", "9999999999999")
	user2, _ := domain.NewUser("mocked2@user.com", "!@#123Mockedpassword", "mocker user 2", "9999999999999")

	repo.Users = append(repo.Users, *user1, *user2)
	return &repo
}

func (ur UserRepositoryMemory) BeginTransaction() (tx Transaction, err error) {
	return MockedTransaction{}, nil
}

func (ur *UserRepositoryMemory) CreateUser(user *domain.User) (err error) {
	ur.Users = append(ur.Users, *user)
	return nil
}

func (ur *UserRepositoryMemory) GetUser(userID *string) (user *domain.User, err error) {
	for i := range ur.Users {
		if ur.Users[i].Id() == *userID {
			return &ur.Users[i], nil
		}
	}
	return user, ops.NewErro("user not found")
}

func (ur *UserRepositoryMemory) ListUsers() (users []domain.User, err error) {
	return ur.Users, nil
}

func (ur *UserRepositoryMemory) UpdateUser(userID *string, user *domain.User) (err error) {
	for i := range ur.Users {
		if ur.Users[i].Id() == *userID {
			ur.Users[i] = *user
			return
		}
	}
	return ops.NewErro("user not found")
}

func (ur *UserRepositoryMemory) DeleteUser(userID *string) (err error) {
	for i := range ur.Users {
		if ur.Users[i].Id() == *userID {
			ur.Users = append(ur.Users[:i], ur.Users[i+1:]...)
			return
		}
	}
	return ops.NewErro("user not found")

}
