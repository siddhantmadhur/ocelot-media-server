package storage

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateConn() (*gorm.DB, error) {
	persistentDir := os.Getenv("PERSISTENT_DIR")
	dbPath := path.Clean(persistentDir)
	info, err := os.Stat(dbPath)
	if err != nil || !info.IsDir() {
		if os.IsNotExist(err) == true {
			err := os.MkdirAll(dbPath, os.ModePerm)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Error: %s\n", err.Error()))
			}
		} else {
			return nil, errors.New(fmt.Sprintf("PERSISTENT_DIR is not valid: (%s)\n Error: %s\n", persistentDir, err.Error()))

		}
	}
	log.Printf("[DATABASE] Connecting to \"%s/storage.db\"...\n", dbPath)
	tx, err := gorm.Open(sqlite.Open(dbPath + "/storage.db"))
	return tx, err
}
