package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	mapURLs()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
