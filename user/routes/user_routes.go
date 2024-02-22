package routes

import (
	"user/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controller.CreateUser())
	router.GET("/user/:userId", controller.GetUserById())
	router.PUT("/user/:userId", controller.UpdateUserById())

}
