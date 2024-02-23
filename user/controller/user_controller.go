package controller

import (
	"net/http"

	"github.com/shublakhan-kaur/job-portal/user/model"
	"github.com/shublakhan-kaur/job-portal/user/service"

	"github.com/gin-gonic/gin"
)

func CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}
		newUser := model.User{
			UserId:    user.UserId,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Work_auth: user.Work_auth,
		}
		result := service.CreateUser(&newUser)
		ctx.JSON(http.StatusCreated, model.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result.InsertedID}})
	}
}

func GetUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("userId")
		var user model.User
		result := service.GetUserById(userId).Decode(&user)
		if result != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": result.Error()}})
		} else {
			ctx.JSON(http.StatusOK, model.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
		}
	}
}

func UpdateUserById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("userId")
		var user model.User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}
		updateUser := model.User{
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Work_auth: user.Work_auth,
		}
		result := service.UpdateUserById(&updateUser, userId)
		if result != nil {
			err := result.Decode(&user)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, model.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			} else {
				ctx.JSON(http.StatusOK, model.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
			}
		}
	}
}
