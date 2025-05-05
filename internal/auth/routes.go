package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/siddhantmadhur/ocelot-media-server/internal/storage"
)

type ReturnData map[string]any

type createUserRouteParam struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Password    string `json:"password"`
}

func CreateUserRoute(c echo.Context) error {

	tx, err := storage.CreateConn()
	if err != nil {
		return c.JSON(500, ReturnData{
			"error":   err.Error(),
			"message": "There was an internal error in opening a connection to the internal database",
		})
	}

	var newUser createUserRouteParam
	c.Bind(&newUser)

	var users int64
	res := tx.Model(&User{}).Count(&users)
	if res.Error != nil {
		return c.JSON(500, ReturnData{
			"error":   res.Error.Error(),
			"message": "There was an error in contacting the internal database",
		})
	}

	userCreated, err := CreateUser(newUser.Username, newUser.DisplayName, newUser.Password)

	if users == 0 {
		// Create root user
		tx.Create(&userCreated)
		return c.JSON(201, ReturnData{
			"message": "Created user!",
		})
	} else {
		// check if allowed and create sub user
	}

	return c.JSON(201, ReturnData{
		"message": "Created user!",
	})
}
