package main

import (
	"flag"
	"fmt"
	"os"

	"git.reocelot.com/ocelot/internal"
	"git.reocelot.com/ocelot/pkg/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Handle flags
	var printVersion = flag.Bool("v", false, "print software version")
	var webServerPort = flag.Int("p", 8080, "web server port to receive incoming traffic")

	flag.Parse()

	if *printVersion {
		fmt.Printf("%s\n", internal.Version)
		os.Exit(0)
	}

	// Start web server

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	api.RegisterRoute(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *webServerPort)))

	os.Exit(0)
}
