package services

import (
	"fmt"
	"github.com/maurobastasini/go-microservices/mvc/model"
	"github.com/maurobastasini/go-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	getUserFunc func(userID uint64) (*model.User, *utils.ApiError)
)

type (
	userDAOMock struct {}
)

func init() {
	model.UserDAO = &userDAOMock{}
}

func(m *userDAOMock) GetUser(userID uint64) (*model.User, *utils.ApiError) {
	return getUserFunc(userID)
}

func TestGetUser_Successful(t *testing.T) {
	getUserFunc = func(userID uint64) (*model.User, *utils.ApiError) {
		return &model.User{
			Id:        userID,
			FirstName: "Mauro",
			LastName:  "Bastasini",
			Email:     "maurobastasiniprof@gmail.com",
		}, nil
	}

	user, err := UsersService.GetUser(1)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Mauro", user.FirstName)
	assert.EqualValues(t, "Bastasini", user.LastName)
	assert.EqualValues(t, "maurobastasiniprof@gmail.com", user.Email)
}

func TestGetUser_NotFound(t *testing.T) {
	getUserFunc = func(userID uint64) (*model.User, *utils.ApiError) {
		return nil, &utils.ApiError{
			Message:    fmt.Sprintf("userID: %d not found", userID),
			StatusCode: http.StatusNotFound,
			Status:     "not_found",
		}
	}

	user, err := UsersService.GetUser(0)
	assert.NotNil(t, err)
	assert.EqualValues(t, "userID: 0 not found", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.Nil(t, user)
}