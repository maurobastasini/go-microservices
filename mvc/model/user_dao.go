package model

import (
	"fmt"
	"github.com/maurobastasini/go-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[uint64]*User{
		1: {
			Id:        1,
			FirstName: "Mauro",
			LastName:  "Bastasini",
			Email:     "maurobastasiniprof@gmail.com",
		},
		2: {
			Id:        2,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@gmail.com",
		},
	}

	UserDAO IUserDAO
)

type (
	IUserDAO interface {
		GetUser(userID uint64) (*User, *utils.ApiError)
	}

	userDAO struct {}
)

func init() {
	UserDAO = &userDAO{}
}

func (u *userDAO) GetUser(userID uint64) (*User, *utils.ApiError) {
	user, found := users[userID]
	if !found {
		return nil, &utils.ApiError{
			Message:    fmt.Sprintf("userID: %d not found", userID),
			StatusCode: http.StatusNotFound,
			Status:     "not_found",
		}
	}

	return user, nil
}
