package user

import "time"

type UserInfoModel struct {
	PK        string
	SK        string
	Email     string
	Name      string
	BirthDate *time.Time
}

type UserLogin struct {
	PK       string
	SK       string
	Password string
	UserID   string
}
