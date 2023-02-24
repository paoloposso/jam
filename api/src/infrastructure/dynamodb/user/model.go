package user

import "time"

type UserModel struct {
	PK        string
	SK        string
	Email     string
	Name      string
	BirthDate *time.Time
}
