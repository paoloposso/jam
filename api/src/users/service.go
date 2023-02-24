package users

import (
	errors "github.com/paoloposso/jam/src/core/custom-errors"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service Service) InsertUser(user User) error {
	if user.Email == "" {
		return errors.CreateValidationError("email is required")
	}
	if user.Name == "" {
		return errors.CreateValidationError("name is required")
	}
	return service.repository.Insert(user)
}
