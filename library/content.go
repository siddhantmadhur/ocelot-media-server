package library

import (
	"context"
	"ocelot/config"
	"ocelot/storage"
	"strconv"

	"github.com/labstack/echo/v4"
)

// /media/library/series | movie | season/content?library=id
func GetContentFromLibrary(c echo.Context, cfg *config.Config) error {
	libraryId, err := strconv.Atoi(c.QueryParam("library"))
	mediaType := c.Param("mediaType")
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(500, result)
	}

	conn, query, err := storage.GetConn(cfg)
	defer conn.Close()
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(500, result)
	}

	contentFiles, err := query.GetAllShows(context.Background(), storage.GetAllShowsParams{
		MediaLibraryID: int64(libraryId),
		MediaType:      mediaType,
	})

	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(500, result)
	}

	if mediaType == "series" {
		var root = []map[string]any{}

		for _, contentFile := range contentFiles {
			root = append(root, map[string]any{
				"title":                contentFile.MediaTitle,
				"description":          contentFile.Description.String,
				"cover_url":            contentFile.CoverUrl.String,
				"media_type":           contentFile.MediaType,
				"external_provider_id": contentFile.ExternalProviderID.Int64,
				"external_provider":    contentFile.ExternalProvider.String,
			})
		}

		return c.JSON(200, root)
	}

	return c.NoContent(500)
}
