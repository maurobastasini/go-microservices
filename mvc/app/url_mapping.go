package app

import "github.com/maurobastasini/go-microservices/mvc/controllers"

func mapURLs() {
	router.GET("/users/:user_id", controllers.GetUser)
}