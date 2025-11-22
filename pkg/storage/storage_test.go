package storage

import (
	"log"
	"testing"
)

func TestConnection(t *testing.T) {
	_, err := CreateConn()
	if err != nil {
		log.Printf("COULD NOT CONNECT TO DATABSE: %s\n", err.Error())
		t.FailNow()
	}
}
