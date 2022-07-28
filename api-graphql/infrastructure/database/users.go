package database

import (
	"api-graphql/users"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Client
}

func (this UserRepository) Insert(user users.User) (id string, err error) {
	println("user inserted")
	return "aaa", nil
}

func (this UserRepository) Update(user users.User) error {
	println("user updated")
	return nil
}
