package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func CreateConn() (*pgx.Conn, error) {
	log.Printf("INFO Creating new database connection")

	ctx := context.Background()

	var (
		host     = os.Getenv("PG_HOST")
		port     = os.Getenv("PG_PORT")
		user     = os.Getenv("PG_USER")
		database = os.Getenv("PG_DATABASE")
		password = os.Getenv("PG_PASSWORD")
	)
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", host, user, database, port, password)

	conn, err := pgx.Connect(ctx, dsn)

	return conn, err
}
