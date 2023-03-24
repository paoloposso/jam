package auth

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

func (srv Service) Authenticate(email string, password string) (*AuthenticatedUser, error) {
	userID, hashedPassword, err := srv.repository.GetUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if core.ValidatePasswordHash(password, hashedPassword) {
		return &AuthenticatedUser{
				Email:  email,
				UserID: userID,
				Token:  core.GetRandomId()},
			nil
	}

	return nil, customerrors.CreateUnauthorizedError()
}
