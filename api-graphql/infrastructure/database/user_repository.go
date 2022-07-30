package database

import (
	"api-graphql/users"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const database = "jamapp"
const collection = "jamapp"

type UserRepository struct {
	*mongo.Client
}

func (this UserRepository) Insert(user users.User) (id string, err error) {
	col := this.Client.Database(database).Collection(collection)

	result, err := col.InsertOne(context.TODO(), user)

	if err != nil {
		return "", err
	}
	return objectIdToString(*result), nil
}

func (this UserRepository) GetByEmail(email string) (*users.User, error) {
	col := this.Client.Database(database).Collection(collection)
	filter := bson.D{{Key: "email", Value: email}}
	var result users.User

	err := col.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (this UserRepository) Update(user users.User) error {
	col := this.Client.Database(database).Collection(collection)

	_, err := col.UpdateByID(context.TODO(), user.ID, user)

	return err
}
