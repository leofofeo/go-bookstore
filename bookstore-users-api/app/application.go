package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

// StartApplication starts the application by calling the router's mapped urls and starting the server on point 8080
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
