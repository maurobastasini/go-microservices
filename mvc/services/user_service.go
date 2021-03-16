package services

import (
	"github.com/maurobastasini/go-microservices/mvc/model"
	"github.com/maurobastasini/go-microservices/mvc/utils"
)

var (
	UsersService usersService
)

type (
	usersService struct {}
)

func (u *usersService) GetUser(userID uint64) (*model.User, *utils.ApiError) {
	return model.UserDAO.GetUser(userID)
}
