package app

import (
	"bookstore-users-api/controllers/ping"
	"bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/pong", ping.Pong)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
