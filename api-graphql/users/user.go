package users

type User struct {
	ID        string
	Email     string
	Name      string
	BirthDate string
	Location  Location
}

type Location struct {
}
