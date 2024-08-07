package library

import (
	"context"
	"database/sql"
	"fmt"
	"ocelot/auth"
	"ocelot/config"
	"ocelot/storage"
	"regexp"
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

func GetVideoContentFromMedia(c echo.Context, u *auth.User, cfg *config.Config) error {
	conn, query, err := storage.GetConn(cfg)
	defer conn.Close()
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(500, result)
	}

	mediaId, err := strconv.Atoi(c.Param("mediaId"))
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(500, result)
	}

	contents, err := query.GetContentFromParentId(context.Background(), sql.NullInt64{Int64: int64(mediaId), Valid: true})
	if err != nil {
		var result = map[string]string{
			"message": "There was an error",
			"error":   err.Error(),
		}
		return c.JSON(500, result)
	}

	var results = []map[string]string{}

	season, err := regexp.Compile("S[0-9]+")
	episode, err := regexp.Compile("E[0-9]+")
	getNumber, err := regexp.Compile("[0-9]+")
	for _, content := range contents {
		current := map[string]string{
			"id":                   fmt.Sprint(content.ID),
			"title":                content.MediaTitle,
			"description":          content.Description.String,
			"external_provider":    content.ExternalProvider.String,
			"external_provider_id": fmt.Sprint(content.ExternalProviderID.Int64),
		}
		if content.MediaType == "episode" {
			current["season"] = getNumber.FindString(season.FindString(content.Classifier))
			current["episode"] = getNumber.FindString(episode.FindString(content.Classifier))
		}
		results = append(results, current)
	}

	return c.JSON(200, results)
}
