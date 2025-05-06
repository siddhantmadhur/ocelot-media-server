package server

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/siddhantmadhur/ocelot-media-server/internal"
	"github.com/siddhantmadhur/ocelot-media-server/internal/storage"
)

func (s *Server) Healthcheck(c echo.Context) error {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	config, _ := storage.GetSettings()

	return c.JSON(200, map[string]any{
		"message":                "Server is healthy!",
		"hostname":               hostname,
		"uptime":                 time.Since(s.StartTime).String(),
		"version":                internal.Version,
		"finished_configuration": config.General.CompletedSetup,
	})
}
