package app

import (
	"bookstore_users-api/controllers/ping"
	"bookstore_users-api/controllers/user"
)

func MapUrls() {

	router.MaxMultipartMemory = 4 << 20 //max file 4mb

	router.GET("/ping", ping.Ping)
	router.GET("/users/:userID", user.GetUser)
	// router.GET("/users/search", user.SearchUser)
	router.POST("/users", user.CreateUser)
	router.PUT("/users/:userID", user.UpdateUser)
	router.PATCH("/users/:userID", user.PatchEmailUser)
	router.DELETE("/users/:userID", user.DeleteUser)
	router.GET("/internal/users/search", user.Search)

	router.POST("/internal/upload/test", user.Upload)
	router.Static("/files", "./files")
	// router.Static("/images", "./images")

}
