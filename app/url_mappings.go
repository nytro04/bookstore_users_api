package app

import "github.com/nytro04/bookstore_users_api/controllers"

func mapUrls()  {
	router.GET("/ping", controllers.Ping)
}
