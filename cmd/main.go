package main

import (
	"log"

	"github.com/siddhantmadhur/ocelot-media-server/internal/auth"
	"github.com/siddhantmadhur/ocelot-media-server/internal/storage"
)

func main() {
	log.Printf("[MAIN] Starting server...\n")

	tx, err := storage.CreateConn()
	if err != nil {
		log.Fatalf("Error connecting to database: %s\n", err.Error())
	}

	auth.UpdateModels(tx)
}
