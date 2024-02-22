package main

import (
	"user/config"
	"user/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.ConnectDB()
	routes.UserRoute(router)
	router.Run("localhost:8000")
}
