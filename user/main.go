package main

import (
	"github.com/shublakhan-kaur/job-portal/user/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoute(router)
	router.Run("localhost:8000")
}
