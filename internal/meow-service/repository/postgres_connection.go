package repository

import (
	"database/sql"

	"github.com/google/wire"
	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	Database *sql.DB
}

func ProvidePostgresConnection() (*PostgresConnection, func(), error) {
	// TODO: Заменить на реальное
	connectionString := "host=localhost port=5432 user=postgres password=postgres dbname=cybermeowerdb sslmode=disable"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err = db.Close(); err != nil {
			// TODO: Добавить обработку ошибок
		}
	}

	return &PostgresConnection{Database: db}, cleanup, nil
}

var PostgresConnectionSet = wire.NewSet(
	ProvidePostgresConnection,
)
