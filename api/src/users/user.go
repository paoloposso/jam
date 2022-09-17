package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     string
	Name      string
	BirthDate string
	Location  Location
}

type Location struct {
	Latitude  float64
	Longitude float64
}
