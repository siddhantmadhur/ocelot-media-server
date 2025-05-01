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
	)

	flag.BoolVar(&printVersion, "v", false, "print software version")
	flag.StringVar(&persistentDir, "p", persistentDir, "change persistent directory")

	flag.Parse()

	if len(os.Getenv("PERSISTENT_DIR")) == 0 {
		os.Setenv("PERSISTENT_DIR", persistentDir)
	}

	if printVersion {
		fmt.Printf("%s\nversion: %s\n", os.Args[0], internal.Version)
	} else {
		log.Printf("[MAIN] Starting server...\n")
		tx, err := storage.CreateConn()
		if err != nil {
			log.Fatalf("Error connecting to database: %s\n", err.Error())
		}

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
