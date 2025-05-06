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

	settings, err := storage.GetSettings()
	if err != nil {
		return c.JSON(500, ReturnData{
			"error":   err.Error(),
			"message": "There was an error in reading the internal settings file",
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
		if !settings.General.CompletedSetup {
			var prevUser User
			tx.First(&prevUser)
			res := tx.Model(&prevUser).Updates(userCreated)
			if res.Error != nil {
				return c.JSON(500, map[string]string{
					"error": res.Error.Error(),
				})
			}
			return c.JSON(200, map[string]string{
				"message": "Updated user!",
			})
		} else {
			// TODO: Create a second user if the route is authorized
		}
	}

	return c.JSON(201, ReturnData{
		"message": "Created user!",
	})
}
