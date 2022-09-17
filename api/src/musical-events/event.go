package musicalevents

import "go.mongodb.org/mongo-driver/bson/primitive"

type MusicalEvent struct {
	Location    Location
	Name        string
	Description string
	ID          primitive.ObjectID
	CreatorID   primitive.ObjectID
}

type Location struct {
	Latitude  float64
	Longitude float64
}
