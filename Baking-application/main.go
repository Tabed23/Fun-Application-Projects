package main

import (
	"log"

	"github.com/Tabed23/Fun-Application-Projects/tree/main/Baking-application/api"
	"github.com/Tabed23/Fun-Application-Projects/tree/main/Baking-application/repository"
)

func main() {

	store, err := repository.NewPostgresStore()

	if err != nil {

		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8000", store)

	server.Run()
}
