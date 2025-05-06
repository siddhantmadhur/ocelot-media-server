package storage

import (
	"fmt"
	"os"
	"testing"
)

func TestSettings(t *testing.T) {
	tempDir := t.TempDir()
	t.Setenv("PERSISTENT_DIR", tempDir)
	_, err := GetSettings()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		t.FailNow()
	}
	data, err := os.ReadFile(tempDir + "/settings.toml")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		t.FailNow()
	}
	fmt.Printf("Data: %s\n", string(data))
}
