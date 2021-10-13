package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/maurobastasini/go-microservices/mvc/services"
	"github.com/maurobastasini/go-microservices/mvc/utils"
)

func GetUser(c *gin.Context) {
	userIdParam, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id param must be a positive number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiError)
		return
	}

	user, apiErr := services.GetUser(userIdParam)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}