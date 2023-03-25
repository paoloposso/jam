package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/paoloposso/jam/libs/auth"
	authrepo "github.com/paoloposso/jam/libs/infrastructure/dynamodb/auth"
	userrepo "github.com/paoloposso/jam/libs/infrastructure/dynamodb/user"
	"github.com/paoloposso/jam/libs/users"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		return nil
	})

	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	repo, err := userrepo.NewUserRepository(svc)
	authRepo, err := authrepo.NewRepository(svc)

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
