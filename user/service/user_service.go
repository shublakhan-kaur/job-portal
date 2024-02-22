package service

import (
	"context"
	"user/config"
	"user/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user *model.User) (*mongo.InsertOneResult, error) {
	userCollection := config.GetCollection(config.DB, config.EnvMongoCollection("MONGO_USER_COLLECTION"))
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	//fmt.Println(result)
	return result, nil

}

func GetUserById(userId string) *mongo.SingleResult {
	userCollection := config.GetCollection(config.DB, config.EnvMongoCollection("MONGO_USER_COLLECTION"))
	result := userCollection.FindOne(context.Background(), bson.M{"userid": userId})
	return result
}

func UpdateUserById(user *model.User, userId string) (*mongo.SingleResult, error) {
	userCollection := config.GetCollection(config.DB, config.EnvMongoCollection("MONGO_USER_COLLECTION"))
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
