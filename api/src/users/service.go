package users

import (
	"github.com/paoloposso/jam/src/core"
	errors "github.com/paoloposso/jam/src/core/custom_errors"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (service Service) CreateUser(user User) error {
	err := validateUser(user)
	if err != nil {
		return err
	}
	user.ID = core.GetRandomId()
	return service.repository.Insert(user)
}

func validateUser(user User) error {
	if user.Email == "" {
		return errors.CreateValidationError("Email is required")
	}
	if user.Name == "" {
		return errors.CreateValidationError("Name is required")
	}
	if user.Password == "" {
		return errors.CreateValidationError("Password is required")
	}
	return nil
}
