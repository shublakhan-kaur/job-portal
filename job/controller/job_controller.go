package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shublakhan-kaur/job-portal/job/model"
	"github.com/shublakhan-kaur/job-portal/job/service"
)

func CreateJob() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var job model.Job
		if err := ctx.BindJSON(&job); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}
		newJob := model.Job{
			JobId:   job.JobId,
			JobDesc: job.JobDesc,
			JobReq:  job.JobReq,
		}
		result := service.CreateJob(&newJob)
		ctx.JSON(http.StatusCreated, model.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result.InsertedID}})
	}
}

func GetJobById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jobId := ctx.Param("jobId")
		var job model.Job
		result := service.GetJobById(jobId).Decode(&job)
		if result != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": result.Error()}})
		} else {
			ctx.JSON(http.StatusOK, model.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": job}})
		}
	}
}

func GetJobs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result := service.GetJobs()
		ctx.JSON(http.StatusOK, model.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
