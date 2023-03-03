package users

import "time"

type User struct {
	ID        string
	Email     string
	Name      string
	BirthDate *time.Time
	Password  string
}
