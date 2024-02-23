package service

import (
	"context"

	"github.com/shublakhan-kaur/job-portal/job/config"
	"github.com/shublakhan-kaur/job-portal/job/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateJob(job *model.Job) *mongo.InsertOneResult {
	var DB *mongo.Client = config.ConnectDB()
	jobCollection := config.GetCollection(DB, config.EnvMongoCollection())
	defer DB.Disconnect(context.Background())
	result, err := jobCollection.InsertOne(context.Background(), job)
	if err != nil {
		panic(err)
	}
	return result
}

func GetJobById(jobId string) *mongo.SingleResult {
	var DB *mongo.Client = config.ConnectDB()
	jobCollection := config.GetCollection(DB, config.EnvMongoCollection())
	defer DB.Disconnect(context.Background())
	result := jobCollection.FindOne(context.Background(), bson.M{"jobid": jobId})
	return result
}

func GetJobs() []model.Job {
	var DB *mongo.Client = config.ConnectDB()
	jobCollection := config.GetCollection(DB, config.EnvMongoCollection())
	defer DB.Disconnect(context.Background())
	result, err := jobCollection.Find(context.Background(), bson.M{})
	var jobs []model.Job
	if err != nil {
		panic(err)
	} else {
		for result.Next(context.Background()) {
			var job model.Job
			err := result.Decode(&job)
			if err != nil {
				panic(err)
			}
			jobs = append(jobs, job)
		}
	}
	return jobs
}
