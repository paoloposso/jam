package users

import (
	"errors"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{repository: repository}
}

func (service Service) InsertUser(user User) (id string, err error) {
	if user.Email == "" {
		return "", errors.New("email is required")
	}
	if user.Name == "" {
		return "", errors.New("name is required")
	}
	if user.BirthDate == "" {
		return "", errors.New("birthdate is required")
	}
	return service.repository.Insert(user)
}

func (service Service) GetByEmail(email string) (user *User, err error) {
	if email == "" {
		return nil, errors.New("email is required")
	}
	return service.repository.GetByEmail(email)
}

func (service Service) GetById(id string) (user *User, err error) {
	if id == "" {
		return nil, errors.New("email is required")
	}
	return service.repository.GetByEmail(id)
}
