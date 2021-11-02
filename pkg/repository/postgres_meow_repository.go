package repository

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/domain"
	"github.com/pkg/errors"
)

type PostgresMeowRepository struct {
	connection *PostgresConnection
}

func ProvidePostgresMeowRepository(connection *PostgresConnection) *PostgresMeowRepository {
	return &PostgresMeowRepository{
		connection: connection,
	}
}

func (repository *PostgresMeowRepository) Create(meow *domain.Meow) (int, error) {
	var id int
	err := repository.connection.database.QueryRow("INSERT INTO meows(body, created_on) VALUES($1, $2) RETURNING id", meow.Body, meow.CreatedOn).Scan(&id)

	if err != nil {
		return 0, errors.Wrap(err, "postgres meow repository")
	}

	return int(id), nil
}

var PostgresMeowRepositorySet = wire.NewSet(
	ProvidePostgresMeowRepository,
	wire.Bind(new(MeowRepository), new(*PostgresMeowRepository)),
)
