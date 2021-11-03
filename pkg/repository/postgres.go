package repository

import (
	"database/sql"

	"github.com/kshvyryaev/cyber-meower-meower-service/pkg"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func ProvidePostgres(config *pkg.Config, logger *zap.Logger) (*sql.DB, func(), error) {
	database, err := sql.Open("postgres", config.DatabaseConnectionString)
	if err != nil {
		return nil, nil, errors.Wrap(err, "postgres database")
	}

	cleanup := func() {
		if err = database.Close(); err != nil {
			logger.Error("cannot close database: " + err.Error())
		}
	}

	return database, cleanup, nil
}
