package app

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// StartApp is a function that manage the api routes
func StartApp() {
	mapUrls()
	router.Run(":8081")
}
