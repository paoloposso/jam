package auth_test

import (
	"fmt"
	"testing"

	"github.com/paoloposso/jam/src/auth"
	"github.com/paoloposso/jam/src/core"
	"github.com/paoloposso/jam/src/core/customerrors"
)

func TestLogin(t *testing.T) {
	repo := AuthRepoMock{}
	service := auth.NewService(repo)
	_, err := service.Authenticate("pvictorsys@gmail.com", "1234")

	if err != nil {
		t.Fatal(err)
	}
}

type AuthRepoMock struct {
}

// GetUserByEmail implements auth.Repository
func (AuthRepoMock) GetUserByEmail(email string) (userId string, password string, err error) {
	if email == "pvictorsys@gmail.com" {
		hashedPassword, err := core.HashPassword("1234")
		if err != nil {
			return "", "", nil
		}
		return "ABC123", hashedPassword, nil
	}
	return "", "", customerrors.CreateArgumentError(fmt.Sprintf("%s not found", email))
}
