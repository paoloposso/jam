package users

import "time"

type User struct {
	ID        string
	Email     string
	Name      string
	Location  Location
	BirthDate *time.Time
}

type Location struct {
	Latitude  float64
	Longitude float64
}
