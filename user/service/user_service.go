package service

import (
	"context"

	"github.com/shublakhan-kaur/job-portal/user/config"
	"github.com/shublakhan-kaur/job-portal/user/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user *model.User) (*mongo.InsertOneResult, error) {
	var DB *mongo.Client = config.ConnectDB()
	userCollection := config.GetCollection(DB, config.EnvMongoCollection("MONGO_USER_COLLECTION"))
	defer DB.Disconnect(context.Background())
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	//fmt.Println(result)

	return result, nil

}

func GetUserById(userId string) *mongo.SingleResult {
	var DB *mongo.Client = config.ConnectDB()
	userCollection := config.GetCollection(DB, config.EnvMongoCollection("MONGO_USER_COLLECTION"))
	defer DB.Disconnect(context.Background())
	result := userCollection.FindOne(context.Background(), bson.M{"userid": userId})
	return result
}

func UpdateUserById(user *model.User, userId string) (*mongo.SingleResult, error) {
	var DB *mongo.Client = config.ConnectDB()
	userCollection := config.GetCollection(DB, config.EnvMongoCollection("MONGO_USER_COLLECTION"))
	defer DB.Disconnect(context.Background())
	update := bson.M{"name": user.Name, "email": user.Email, "phone": user.Phone, "work_auth": user.Work_auth}
	result, err := userCollection.UpdateOne(context.Background(), bson.M{"userid": userId}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 1 {
		return GetUserById(userId), nil
	}
	return nil, nil
}
