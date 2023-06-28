package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/severusTI/auth_golang/internal/interfaces/api/handlers"
	"github.com/severusTI/auth_golang/internal/interfaces/api/routers"
	"github.com/severusTI/auth_golang/internal/interfaces/persistance/repositories"
	"github.com/severusTI/auth_golang/internal/usecases"
	"github.com/severusTI/auth_golang/internal/usecases/dtos"
	"github.com/severusTI/auth_golang/pkg/database"
	env "github.com/severusTI/auth_golang/pkg/env_load"
	"github.com/stretchr/testify/assert"
)

func TestUsersEndpoints(t *testing.T) {
	env.LoadEnv()
	db := database.InitDBTESTConnection()
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userusecase := usecases.NewUserUC(userRepo)
	userHandlers := handlers.NewUserHandlers(*userusecase)

	router := routers.SetupUserRoutes(*userHandlers)

	// TESTING CREATE USER
	payload := map[string]string{
		"name":         "user test 1",
		"email":        "test1@test.com",
		"phone_number": "5585999999999",
		"password":     "123!@#Password",
	}
	payloadBytes, _ := json.Marshal(payload)

	payload2 := map[string]string{
		"name":         "user test 2",
		"email":        "test2@test.com",
		"phone_number": "5585999999998",
		"password":     "123!@#Password",
	}
	payloadBytes2, _ := json.Marshal(payload2)

	// Cria uma requisição HTTP POST para o endpoint /users
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewReader(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	req2, _ := http.NewRequest("POST", "/api/users", bytes.NewReader(payloadBytes2))
	req2.Header.Set("Content-Type", "application/json")

	// Cria um ResponseRecorder para gravar a resposta do servidor
	res := httptest.NewRecorder()
	res2 := httptest.NewRecorder()

	router.ServeHTTP(res, req)
	router.ServeHTTP(res2, req2)
	assert.Equal(t, http.StatusCreated, res.Code)
	assert.Equal(t, http.StatusCreated, res2.Code)

	// TESTING LISTING USERS
	req3, _ := http.NewRequest("GET", "/api/users/list", nil)
	res3 := httptest.NewRecorder()

	router.ServeHTTP(res3, req3)
	assert.Equal(t, http.StatusOK, res3.Code)
	var responseUsers = make([]dtos.ResUser, 0)
	err := json.Unmarshal(res3.Body.Bytes(), &responseUsers)
	assert.NoError(t, err)
	assert.Len(t, responseUsers, 2)
	assert.Equal(t, "user test 1", responseUsers[0].Name)
	assert.Equal(t, "user test 2", responseUsers[1].Name)
	assert.Equal(t, "test1@test.com", responseUsers[0].Email)
	assert.Equal(t, "test2@test.com", responseUsers[1].Email)
	assert.Equal(t, "5585999999999", responseUsers[0].PhoneNumber)
	assert.Equal(t, "5585999999998", responseUsers[1].PhoneNumber)

	// TESTING GET USER
	req4, _ := http.NewRequest("GET", "/api/users/"+responseUsers[0].ID, nil)
	res4 := httptest.NewRecorder()
	router.ServeHTTP(res4, req4)
	var responseUser dtos.ResUser
	err = json.Unmarshal(res4.Body.Bytes(), &responseUser)
	assert.NoError(t, err)
	assert.Equal(t, "user test 1", responseUser.Name)
	assert.Equal(t, "test1@test.com", responseUser.Email)
	assert.Equal(t, "5585999999999", responseUser.PhoneNumber)

	// TEST UPDATING USER
	payload3 := map[string]string{
		"name":         "user updated",
		"email":        "emailupdated@test.com",
		"phone_number": "5585999999988",
		"password":     "123!@#Updatepassword",
	}
	payloadBytes3, _ := json.Marshal(payload3)

	req5, _ := http.NewRequest("PUT", "/api/users/"+responseUsers[0].ID, bytes.NewReader(payloadBytes3))
	res5 := httptest.NewRecorder()
	router.ServeHTTP(res5, req5)
	assert.Equal(t, http.StatusNoContent, res5.Code)
	req6, _ := http.NewRequest("GET", "/api/users/"+responseUsers[0].ID, nil)
	res6 := httptest.NewRecorder()
	router.ServeHTTP(res6, req6)
	var responseUser2 dtos.ResUser
	err = json.Unmarshal(res6.Body.Bytes(), &responseUser2)
	assert.NoError(t, err)
	assert.Equal(t, "user updated", responseUser2.Name)
	assert.Equal(t, "emailupdated@test.com", responseUser2.Email)
	assert.Equal(t, "5585999999988", responseUser2.PhoneNumber)

	// TEST DELETING USER
	req7, _ := http.NewRequest("DELETE", "/api/users/"+responseUsers[0].ID, bytes.NewReader(payloadBytes3))
	res7 := httptest.NewRecorder()
	router.ServeHTTP(res7, req7)
	assert.Equal(t, http.StatusNoContent, res5.Code)
	req8, _ := http.NewRequest("GET", "/api/users/list", nil)
	res8 := httptest.NewRecorder()

	router.ServeHTTP(res8, req8)
	assert.Equal(t, http.StatusOK, res8.Code)
	var responseUsers2 = make([]dtos.ResUser, 0)
	err = json.Unmarshal(res8.Body.Bytes(), &responseUsers2)
	assert.NoError(t, err)
	assert.Len(t, responseUsers2, 1)

}
