package users

type User struct {
	ID        string `bson:"_id"`
	Email     string
	Name      string
	BirthDate string
	Location  Location
}

type Location struct {
	Latitude  float64
	Longitude float64
}
