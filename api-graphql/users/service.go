package users

import "errors"

type Service struct {
	Repository Repository
}

func (this Service) InsertUser(user User) (id string, err error) {
	if user.Email == "" {
		return "", errors.New("email is required")
	}
	if user.Name == "" {
		return "", errors.New("name is required")
	}
	if user.BirthDate == "" {
		return "", errors.New("birthdate is required")
	}
	return this.Repository.Insert(user)
}
