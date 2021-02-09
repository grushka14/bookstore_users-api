package app

import "github.com/grushka14/bookstore_users-api/src/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.POST("/user", controllers.CreateUser)
	router.GET("/user", controllers.GetUsers)
}
