package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shublakhan-kaur/job-portal/job/config"
	"github.com/shublakhan-kaur/job-portal/job/routes"
)

func main() {
	router := gin.Default()
	routes.JobRoute(router)
	router.Run(config.EnvRouterURI())
}
