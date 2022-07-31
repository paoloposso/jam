package users

import (
	"errors"
	"strings"
	"testing"
)

func TestGetUserByEmail(t *testing.T) {
	service := NewService(&RepositoryMock{})

	user, err := service.GetByEmail("pvictorsys@gmail.com")

	if err != nil {
		t.Error(err)
	}
	if user == nil || user.Email != "pvictorsys@gmail.com" {
		t.Error("Incorrect user returned")
	}
}

func TestGetUserByEmailWithEmailBlank(t *testing.T) {
	service := NewService(&RepositoryMock{})

	_, err := service.GetByEmail("")

	if err == nil {
		t.Fail()
	}
}

func TestGetNilWhenEmailDoesntExist(t *testing.T) {
	service := NewService(&RepositoryMock{})

	user, _ := service.GetByEmail("inexistent@gmail.com")

	if user != nil {
		t.Fail()
	}
}

func TestInsertWhenEmailIsBlank(t *testing.T) {
	service := NewService(&RepositoryMock{})

	_, err := service.InsertUser(User{Email: "", Name: "", BirthDate: ""})

	if err == nil || !strings.Contains(err.Error(), "required") {
		t.Fail()
	}
}

type RepositoryMock struct {
}

func (this *RepositoryMock) Insert(user User) (id string, err error) {
	return "", nil
}

func (this *RepositoryMock) Update(user User) error {
	return errors.New("not implemented")
}

func (this *RepositoryMock) GetByEmail(email string) (*User, error) {

	usersList := []User{
		{Email: "pvictorsys@gmail.com", ID: "ASD1234", Name: "Paolo"},
	}

	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == email {
			return &usersList[i], nil
		}
	}

	return nil, nil
}
