package api

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) {
	// v1
	v1 := e.Group("v1")

	v1.GET("/healthcheck", healthcheck)
}
