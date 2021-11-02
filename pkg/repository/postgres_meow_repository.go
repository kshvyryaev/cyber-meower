package repository

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/domain"
	"github.com/pkg/errors"
)

type PostgresMeowRepository struct {
	database *sql.DB
}

func ProvidePostgresMeowRepository(database *sql.DB) *PostgresMeowRepository {
	return &PostgresMeowRepository{
		database: database,
	}
}

func (repository *PostgresMeowRepository) Create(meow *domain.Meow) (int, error) {
	var id int
	err := repository.database.QueryRow("INSERT INTO meows(body, created_on) VALUES($1, $2) RETURNING id", meow.Body, meow.CreatedOn).Scan(&id)

	if err != nil {
		return 0, errors.Wrap(err, "postgres meow repository")
	}

	return int(id), nil
}

var PostgresMeowRepositorySet = wire.NewSet(
	ProvidePostgresMeowRepository,
	wire.Bind(new(MeowRepository), new(*PostgresMeowRepository)),
)
