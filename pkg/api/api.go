package api

import (
	"context"

	"git.siddhantmadhur.com/ocelot/pkg/storage"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo) {
	// v1
	v1 := e.Group("v1")

	v1.GET("/healthcheck", func(c echo.Context) error {
		ctx := context.Background()
		tx, err := storage.CreateConn()
		defer tx.Close(ctx)
		if err != nil {
			return c.JSON(500, map[string]string{
				"error": "database not healthy!",
			})
		}
		return c.JSON(200, map[string]string{
			"healthy": "ok",
		})
	})
}
