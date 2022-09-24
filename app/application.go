package app

import (
	"bookstore_users-api/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Use(cors.Default())

	MapUrls()

	logger.Info("about to start the application...")
	router.Run(":9393")
}
