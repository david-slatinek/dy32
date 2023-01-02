package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func IDToString(id *mongo.InsertOneResult) string {
	return id.InsertedID.(primitive.ObjectID).String()
}
