package users

import (
	"errors"
	"testing"

	customerrors "github.com/paoloposso/jam/api/src/core/custom-errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	_, ok := err.(*customerrors.ValidationError)
	if err == nil || !ok {
		t.Fatal("Error should be of type ValidationError")
	}
}

func TestGetUserById(t *testing.T) {
	service := NewService(&RepositoryMock{})
	user, err := service.GetById("630558ae2ce333d84361e635")
	if err != nil {
		t.Fail()
	}
	if user == nil || user.Email == "" {
		t.Fail()
	}
}

type RepositoryMock struct {
}

func (r *RepositoryMock) Insert(user User) (id string, err error) {
	return "", nil
}

func (r *RepositoryMock) Update(user User) error {
	return errors.New("not implemented")
}

func (r *RepositoryMock) GetByEmail(email string) (*User, error) {
	usersList := [...]User{
		{Email: "pvictorsys@gmail.com", ID: primitive.NewObjectID(), Name: "Paolo"},
	}
	for i := 0; i < len(usersList); i++ {
		if usersList[i].Email == email {
			return &usersList[i], nil
		}
	}
	return nil, nil
}

func (r *RepositoryMock) GetById(id string) (*User, error) {
	objectId, _ := primitive.ObjectIDFromHex("630558ae2ce333d84361e635")
	usersList := [...]User{
		{Email: "pvictorsys@gmail.com", ID: objectId, Name: "Paolo"},
	}
	for i := 0; i < len(usersList); i++ {
		if usersList[i].ID.Hex() == id {
			return &usersList[i], nil
		}
	}
	return nil, nil
}
