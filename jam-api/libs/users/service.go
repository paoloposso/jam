package users

import (
	"github.com/paoloposso/jam/libs/core"
	"github.com/paoloposso/jam/libs/core/customerrors"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{repository: repository}
}

func (service Service) CreateUser(user User) error {
	err := validateUserInfo(user)
	if err != nil {
		return err
	}
	user.ID = core.GetRandomId()
	hashed, err := core.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashed
	return service.repository.Insert(user)
}

func validateUserInfo(user User) error {
	if user.Email == "" {
		return customerrors.CreateValidationError("Email is required")
	}
	if user.Name == "" {
		return customerrors.CreateValidationError("Name is required")
	}
	if user.Password == "" {
		return customerrors.CreateValidationError("Password is required")
	}
	return nil
}
