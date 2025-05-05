package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/siddhantmadhur/ocelot-media-server/internal/auth"
)

type Server struct {
	Port      int
	Handler   *echo.Echo
	StartTime time.Time
}

func NewServer(port int) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} [SERVER] ${status} - ${method} ${uri}\n",
		CustomTimeFormat: "2006/01/02 15:04:05",
	}))

	var s = Server{
		Port:    port,
		Handler: e,
	}

	// Do routes here
	e.GET("/health", s.Healthcheck)
	e.POST("/auth/user/create", auth.CreateUserRoute)

	return &s
}

func (s *Server) Run() {
	log.Printf("[MAIN] Starting server on port %d\n", s.Port)

	s.StartTime = time.Now()

	go func() {
		s.Handler.Start(fmt.Sprintf(":%d", s.Port))
	}()
}

func (s *Server) Stop() {
	log.Printf("[SERVER] Stopping server...\n")
	s.Handler.Shutdown(context.Background())
}

func (s *Server) Restart() {
	s.Stop()
	s.Run()
}
