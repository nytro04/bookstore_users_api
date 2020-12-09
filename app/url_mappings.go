package app

import (
	"github.com/nytro04/bookstore_users_api/controllers/ping"
	"github.com/nytro04/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// User mappings
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)

	//router.GET("/search", controllers.SearchUser)
}
