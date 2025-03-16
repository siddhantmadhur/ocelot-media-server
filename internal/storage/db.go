package storage

import (
	"errors"
	"log"
	"os"
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateConn() (*gorm.DB, error) {
	persistentDir := os.Getenv("PERSISTENT_DIR")
	info, err := os.Stat(persistentDir)
	if err != nil || !info.IsDir() {
		return nil, errors.New("PERSISTENT_DIR is not valid")
	}
	dbPath := path.Clean(persistentDir)
	log.Printf("[DATABASE] Connecting to \"%s\"...\n", dbPath)
	tx, err := gorm.Open(sqlite.Open(dbPath + "/storage.db"))
	return tx, err
}
