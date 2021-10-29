package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/config"
	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	Database *sql.DB
}

func ProvidePostgresConnection(config *config.Config) (*PostgresConnection, func(), error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Name, config.Database.SSLMode)

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
