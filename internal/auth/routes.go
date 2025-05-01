package auth

import "github.com/labstack/echo/v4"

type ReturnData map[string](map[string]string)

func CreateUserRoute(c echo.Context) error {
	return c.JSON(200, ReturnData{
		"data": map[string]string{
			"user_id":    "",
			"created_at": "",
		},
	})
}
