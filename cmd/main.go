package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/siddhantmadhur/ocelot-media-server/internal"
	"github.com/siddhantmadhur/ocelot-media-server/internal/auth"
	"github.com/siddhantmadhur/ocelot-media-server/internal/server"
	"github.com/siddhantmadhur/ocelot-media-server/internal/storage"
)

func main() {

	var (
		printVersion  bool
		persistentDir string = "/etc/opt/ocelot"
		port          int
	)

	flag.BoolVar(&printVersion, "v", false, "print software version")
	flag.BoolVar(&printVersion, "version", false, "print software version")
	flag.StringVar(&persistentDir, "p", persistentDir, "change persistent directory")
	flag.IntVar(&port, "o", 8080, "change port for the web server")

	flag.Parse()

	if len(os.Getenv("PERSISTENT_DIR")) == 0 {
		os.Setenv("PERSISTENT_DIR", persistentDir)
	}

	if printVersion {
		fmt.Printf("%s\nversion: %s\n", os.Args[0], internal.Version)
	} else {
		log.Printf("[MAIN] Starting server...\n")
		tx, err := storage.GetConnection()
		defer storage.CloseConnection(tx)
		if err != nil {
			log.Fatalf("Error connecting to database: %s\n", err.Error())
		}

		// Update models here
		auth.UpdateModels(tx)

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		s := server.NewServer(8080)
		s.Run()

		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		s.Stop()

	}

}
