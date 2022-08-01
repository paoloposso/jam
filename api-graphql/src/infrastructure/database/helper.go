package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func objectIdToString(result mongo.InsertOneResult) string {
	return result.InsertedID.(primitive.ObjectID).Hex()
}
