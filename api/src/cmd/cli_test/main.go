package main

import (
	"log"

	"github.com/paoloposso/jam/src/infrastructure"
	musicalevents "github.com/paoloposso/jam/src/musical_events"
)

func main() {

	repo, err := infrastructure.NewRepository()

	if err != nil {
		log.Fatalf("Error creating repo: %v", err)
	}

	service := musicalevents.NewService(repo)

	service.CreateEvent(musicalevents.MusicalEvent{})
}
