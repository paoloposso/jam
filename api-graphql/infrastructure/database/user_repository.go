package database

import (
	"api-graphql/users"
	"context"
	"fmt"

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
	return fmt.Sprintf("%b", result.InsertedID), nil
}

func (this UserRepository) Update(user users.User) error {
	return nil
}
