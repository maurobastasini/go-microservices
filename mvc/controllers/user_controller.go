package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maurobastasini/go-microservices/mvc/services"
	"github.com/maurobastasini/go-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userQueryParam := c.Param("user_id")
	userID, err := strconv.ParseUint(userQueryParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApiError{
			Message:    fmt.Sprintf("user_id (%s) must be a number\n%s", userQueryParam, err.Error()),
			StatusCode: http.StatusBadRequest,
			Status:     "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}

	utils.RespondSuccessful(c, user)
}
