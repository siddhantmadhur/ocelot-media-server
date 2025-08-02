package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateConn() (*pgx.Conn, error) {
	ctx := context.Background()

	var (
		host     = "localhost"
		port     = "5432"
		user     = "postgres"
		database = "postgres"
	)
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s", host, user, database, port)

	conn, err := pgx.Connect(ctx, dsn)

	return conn, err
}
