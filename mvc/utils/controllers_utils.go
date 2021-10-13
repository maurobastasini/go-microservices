package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, code int, obj interface{}) {
	if c.GetHeader("Accept") ==  "application/xml" {
		c.XML(code, obj)
	}

	c.JSON(code, obj)
}

func RespondError(c *gin.Context, error *ApplicationError) {
	Respond(c, error.StatusCode, error)
}