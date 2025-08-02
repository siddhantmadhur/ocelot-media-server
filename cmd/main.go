package main

import (
	"git.siddhantmadhur.com/ocelot/pkg/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	api.RegisterRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
