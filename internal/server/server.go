package server

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Port    int
	Handler *echo.Echo
}

func NewServer(port int) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} [SERVER] ${status} - ${method} ${uri}\n",
		CustomTimeFormat: "2006/01/02 15:04:05",
	}))
	return &Server{
		Port:    port,
		Handler: e,
	}
}

func (s *Server) Run() {
	log.Printf("[MAIN] Starting server on port %d\n", s.Port)
	go func() {
		s.Handler.Start(fmt.Sprintf(":%d", s.Port))
	}()
}

func (s *Server) Stop() {
	s.Handler.Shutdown(context.Background())
}

func (s *Server) Restart() {
	s.Stop()
	s.Run()
}
