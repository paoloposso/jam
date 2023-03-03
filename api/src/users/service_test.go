package users_test

import (
	"testing"
	"time"

	"github.com/paoloposso/jam/src/core/customerrors"
	"github.com/paoloposso/jam/src/users"
)

func TestCreateUser(t *testing.T) {
	service := users.NewService(MockUserRepo{})

	tm, _ := time.Parse("2006-01-02", "1988-03-20")

	err := service.CreateUser(users.User{Email: "pvictorsys@gmail.com", Name: "Paolo Test", Password: "999999", BirthDate: &tm})

	if err != nil {
		t.Error(err)
	}
}

func TestCreateUser_ShouldFailUserInvalid(t *testing.T) {
	service := users.NewService(MockUserRepo{})

	tm, _ := time.Parse("2006-01-02", "1988-03-20")

	err := service.CreateUser(users.User{Name: "Paolo Test", Password: "999999", BirthDate: &tm})

	if err == nil {
		t.Fatal("Should Fail for invalid e-mail")
	}

	t.Fail()
}

type MockUserRepo struct {
}

// Get implements users.Repository
func (MockUserRepo) Get(id string) (*users.User, error) {
	if id == "123321" {
		return &users.User{ID: "123321", Name: "Paolo Test", Email: "pvictorsys@gmail.com"}, nil
	}
	return nil, customerrors.CreateArgumentError("User not found")
}

// Insert implements users.Repository
func (MockUserRepo) Insert(users.User) error {
	return nil
}
