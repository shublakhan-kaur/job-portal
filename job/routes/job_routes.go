package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shublakhan-kaur/job-portal/job/controller"
)

func JobRoute(router *gin.Engine) {
	router.POST("/job", controller.CreateJob())
	router.GET("/job/:jobId", controller.GetJobById())
	router.GET("/jobs", controller.GetJobs())

}
