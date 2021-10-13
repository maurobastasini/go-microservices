package services

import (
	"github.com/maurobastasini/go-microservices/mvc/domain"
	"github.com/maurobastasini/go-microservices/mvc/utils"
)

func GetUser(userId uint64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
