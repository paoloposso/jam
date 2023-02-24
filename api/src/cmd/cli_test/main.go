package main

import (
	"fmt"
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
	tm, err := time.Parse("2006-01-02", "1988-02-05")

	err = service.CreateUser(users.User{Email: "pvictorsys@gmail.com", Password: "1234", Name: "Paolo Test", BirthDate: &tm})

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
