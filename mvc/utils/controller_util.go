package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondSuccessful(c *gin.Context, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(http.StatusOK, body)
		return
	}

	c.JSON(http.StatusOK, body)
}

func RespondError(c *gin.Context, apiError *ApiError) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(apiError.StatusCode, apiError)
		return
	}

	c.JSON(apiError.StatusCode, apiError)
}