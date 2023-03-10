package main

import (
	"fmt"
	"log"
	"time"

	"github.com/paoloposso/jam/src/auth"
	authrepo "github.com/paoloposso/jam/src/infrastructure/dynamodb/auth"
	userrepo "github.com/paoloposso/jam/src/infrastructure/dynamodb/user"
	"github.com/paoloposso/jam/src/users"
)

func main() {
	repo, err := userrepo.NewUserRepository()

	authRepo, err := authrepo.NewRepository()

	if err != nil {
		log.Fatalf("Error creating repo: %v", err)
	}

	authService := auth.NewService(authRepo)

	userservice := users.NewService(repo)
	tm, err := time.Parse("2006-01-02", "1988-02-05")

	err = userservice.CreateUser(users.User{Email: "pvictorsys@gmail.com", Password: "1234", Name: "Paolo Test", BirthDate: &tm})

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	authService.Authenticate("pvictorsys@gmail.com", "1234")
}
