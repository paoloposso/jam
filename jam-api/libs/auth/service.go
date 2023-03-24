package auth

import (
	"github.com/paoloposso/jam/libs/core"
	"github.com/paoloposso/jam/libs/core/customerrors"
)

type AuthService struct {
	repository Repository
}

func NewService(repository Repository) AuthService {
	return AuthService{repository: repository}
}

func (auth AuthService) Authenticate(email string, password string) (*AuthenticatedUser, error) {
	userID, hashedPassword, err := auth.repository.GetUserByEmail(email)

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
