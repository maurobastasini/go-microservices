package domain

import (
	"fmt"
	"net/http"

	"github.com/maurobastasini/go-microservices/mvc/utils"
)

var (
	users = map[uint64]*User{
		1: {
			Id:        1,
			FirstName: "Mauro",
			LastName:  "Bastasini",
			Email:     "maurobastasiniprof@gmail.com",
		},
	}
)

func GetUser(userId uint64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user with id %d was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
