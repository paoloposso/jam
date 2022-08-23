package database

import (
	"api-graphql/src/users"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collection = "users"

type UserRepository struct {
	client   *mongo.Client
	database string
}

func NewRepository(url string, databaseName string) *UserRepository {
	client, err := getClient(url)
	if err != nil {
		log.Fatal(err)
	}
	col := client.Database(databaseName).Collection(collection)
	_, err = col.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return &UserRepository{client: client, database: databaseName}
}

func (repo UserRepository) Insert(user users.User) (id string, err error) {
	col := repo.client.Database(repo.database).Collection(collection)
	user.ID = primitive.NewObjectID()
	result, err := col.InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}
	return objectIdToString(*result), nil
}

func (repo UserRepository) GetByEmail(email string) (*users.User, error) {
	col := repo.client.Database(repo.database).Collection(collection)
	filter := bson.D{{Key: "email", Value: email}}
	var result users.User
	err := col.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo UserRepository) GetById(id string) (*users.User, error) {
	col := repo.client.Database(repo.database).Collection(collection)
	filter := bson.D{{Key: "_id", Value: id}}
	var result users.User
	err := col.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo UserRepository) Update(user users.User) error {
	col := repo.client.Database(repo.database).Collection(collection)
	_, err := col.UpdateByID(context.Background(), user.ID, user)
	return err
}
