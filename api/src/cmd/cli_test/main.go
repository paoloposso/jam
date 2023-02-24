package main

import (
	"log"
	"time"

	userrepo "github.com/paoloposso/jam/src/infrastructure/dynamodb/user"
	"github.com/paoloposso/jam/src/users"
)

func main() {
	repo, err := userrepo.NewUserRepository()

	if err != nil {
		log.Fatalf("Error creating repo: %v", err)
	}

	service := users.NewService(repo)

	tm, err := time.Parse(time.RFC3339, "1988-02-05")

	service.InsertUser(users.User{ID: "aaaaa", Email: "pvictorsys@gmail.com", Name: "Paolo Test", BirthDate: &tm})
}
