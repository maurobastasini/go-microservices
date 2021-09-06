package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/maurobastasini/go-microservices/mvc/app/services"
	"github.com/maurobastasini/go-microservices/mvc/utils"
)

func GetUser(writer http.ResponseWriter, request *http.Request) {
	userIdParam, err := strconv.ParseUint(request.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id param must be a positive number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		invalidUserError, _ := json.Marshal(apiError)
		writer.WriteHeader(apiError.StatusCode)
		writer.Write(invalidUserError)
		return
	}

	user, apiErr := services.GetUser(userIdParam)
	if apiErr != nil {
		errJson, _ := json.Marshal(apiErr)
		writer.WriteHeader(apiErr.StatusCode)
		writer.Write(errJson)
		return
	}

	userJson, _ := json.Marshal(user)
	writer.Write(userJson)
}
