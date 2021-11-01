package repository

import (
	"database/sql"

	"github.com/kshvyryaev/cyber-meower/internal/meow-service/config"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type PostgresConnection struct {
	database *sql.DB
}

func ProvidePostgresConnection(config *config.Config, logger *zap.Logger) (*PostgresConnection, func(), error) {
	db, err := sql.Open("postgres", config.DatabaseConnectionString)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err = db.Close(); err != nil {
			logger.Error("cannot close database: " + err.Error())
		}
	}

	return &PostgresConnection{database: db}, cleanup, nil
}
