package app

import (
	"bookstore_users-api/controllers/ping"
	"bookstore_users-api/controllers/user"
)

func MapUrls() {

	router.GET("/ping", ping.Ping)
	router.GET("/users/:userID", user.GetUser)
	// router.GET("/users/search", user.SearchUser)
	router.POST("/users", user.CreateUser)
	router.PUT("/users/:userID", user.UpdateUser)
	router.PATCH("/users/:userID", user.PatchEmailUser)
	router.DELETE("/users/:userID", user.DeleteUser)
	router.GET("/internal/users/search", user.Search)

}
