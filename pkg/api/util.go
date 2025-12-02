package api

import (
	"context"

	"git.reocelot.com/ocelot/pkg/storage"
	"github.com/labstack/echo/v4"
)

func healthcheck(c echo.Context) error {
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
}
