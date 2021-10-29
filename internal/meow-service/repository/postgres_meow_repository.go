package repository

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/domain"
	"github.com/pkg/errors"
)

type PostgresMeowRepository struct {
	Connection *PostgresConnection
}

func (repository *PostgresMeowRepository) Create(meow *domain.Meow) (int, error) {
	var id int
	err := repository.Connection.Database.QueryRow("INSERT INTO meows(body, created_on) VALUES($1, $2) RETURNING id", meow.Body, meow.CreatedOn).Scan(&id)

	if err != nil {
		return 0, errors.Wrap(err, "postgres meow repository error")
	}

	return int(id), nil
}

func ProvidePostgresMeowRepository(connection *PostgresConnection) *PostgresMeowRepository {
	return &PostgresMeowRepository{
		Connection: connection,
	}
}

var PostgresMeowRepositorySet = wire.NewSet(
	ProvidePostgresMeowRepository,
	wire.Bind(new(MeowRepository), new(*PostgresMeowRepository)),
)
