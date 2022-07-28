package users

import "errors"

type Service struct {
	Repository Repository
}

func (this Service) InsertUser(user User) (id string, err error) {
	if user.Email == "" {
		return "", errors.New("email is required")
	}
	return this.Repository.Insert(user)
}
