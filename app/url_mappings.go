package app

import (
	"github.com/nytro04/bookstore_users_api/controllers/ping"
	"github.com/nytro04/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// User mappings
	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)

	//router.GET("/search", controllers.SearchUser)
}
