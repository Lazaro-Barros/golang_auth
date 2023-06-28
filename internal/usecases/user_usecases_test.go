package usecases

import (
	"testing"

	"github.com/severusTI/auth_golang/internal/interfaces/persistance/repositories"
	"github.com/severusTI/auth_golang/internal/usecases/dtos"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Create a mock repository
	repo := repositories.NewUserRepositoryMemory()

	// Create a new instance of UserUC
	uc := NewUserUC(repo)

	// Create a sample request user
	reqUser := &dtos.ReqUser{
		Name:            "John Doe",
		Email:           "johndoe@example.com",
		PhoneNumber:     "1234567890",
		InputedPassword: "!A2password",
	}

	// Call the CreateUser method
	err := uc.CreateUser(reqUser)

	// Assert that no error occurred
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	// Create a mock repository
	repo := repositories.NewUserRepositoryMemory()

	// Create a new instance of UserUC
	uc := NewUserUC(repo)

	// Create a sample user ID
	userID := repo.Users[0].Id()

	// Call the GetUser method
	resUser, err := uc.GetUser(&userID)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the returned user matches the expected user
	assert.Equal(t, repo.Users[0].Id(), resUser.ID)
	assert.Equal(t, repo.Users[0].Name(), resUser.Name)
	assert.Equal(t, repo.Users[0].Email(), resUser.Email)
	assert.Equal(t, repo.Users[0].PhoneNumber(), resUser.PhoneNumber)

}

func TestListUsers(t *testing.T) {
	// Create a mock repository
	repo := repositories.NewUserRepositoryMemory()

	// Create a new instance of UserUC
	uc := NewUserUC(repo)

	// Call the ListUsers method
	resUsers, err := uc.ListUsers()

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the returned users match the expected users
	assert.Len(t, resUsers, len(repo.Users))
	for i := range resUsers {
		assert.Equal(t, repo.Users[i].Id(), resUsers[i].ID)
		assert.Equal(t, repo.Users[i].Name(), resUsers[i].Name)
		assert.Equal(t, repo.Users[i].Email(), resUsers[i].Email)
		assert.Equal(t, repo.Users[i].PhoneNumber(), resUsers[i].PhoneNumber)
	}

}

func TestUpdateUser(t *testing.T) {
	// Create a mock repository
	repo := repositories.NewUserRepositoryMemory()

	// Create a new instance of UserUC
	uc := NewUserUC(repo)

	// Create a sample user ID
	userID := repo.Users[0].Id()

	// Create a sample request user
	reqUser := &dtos.ReqUser{
		Name:            "John Doe",
		Email:           "johndoe@example.com",
		PhoneNumber:     "1234567890",
		InputedPassword: "!1Apassword",
	}

	// Call the UpdateUser method
	err := uc.UpdateUser(userID, reqUser)

	// Assert that no error occurred
	assert.NoError(t, err)

}

func TestDeleteUser(t *testing.T) {
	// Create a mock repository
	repo := repositories.NewUserRepositoryMemory()

	// Create a new instance of UserUC
	uc := NewUserUC(repo)

	// Create a sample user ID
	userID := repo.Users[0].Id()

	// Call the DeleteUser method
	err := uc.DeleteUser(&userID)

	// Assert that no error occurred
	assert.NoError(t, err)

}
