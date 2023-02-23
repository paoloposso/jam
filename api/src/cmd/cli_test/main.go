package main

import (
	"log"

	database "github.com/paoloposso/jam/src/infrastructure/dynamodb"
	musicalevents "github.com/paoloposso/jam/src/musical_events"
)

func main() {
	repo, err := database.NewRepository()

	if err != nil {
		log.Fatalf("Error creating repo: %v", err)
	}

	service := musicalevents.NewService(repo)

	service.CreateEvent(musicalevents.MusicalEvent{})
}
