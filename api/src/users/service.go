package users

import (
	errors "github.com/paoloposso/jam/api/src/core/custom-errors"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service Service) InsertUser(user User) (id string, err error) {
	if user.Email == "" {
		return "", errors.CreateValidationError("email is required")
	}
	if user.Name == "" {
		return "", errors.CreateValidationError("name is required")
	}
	if user.BirthDate == "" {
		return "", errors.CreateValidationError("birthdate is required")
	}
	return service.repository.Insert(user)
}

func (service Service) GetByEmail(email string) (user *User, err error) {
	if email == "" {
		return nil, errors.CreateValidationError("email is required")
	}
	return service.repository.GetByEmail(email)
}

func (service Service) GetById(id string) (user *User, err error) {
	if id == "" {
		return nil, errors.CreateValidationError("email is required")
	}
	return service.repository.GetById(id)
}
