//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/command"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/config"
	controller "github.com/kshvyryaev/cyber-meower/internal/meow-service/controller/http"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/repository"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/service"
)

func InitializeHttpServer() (*controller.HttpServer, func(), error) {
	panic(wire.Build(
		config.ConfigSet,
		repository.PostgresConnectionSet,
		repository.PostgresMeowRepositorySet,
		service.MeowTranslatorServiceSet,
		command.СreateMeowCommandHandlerSet,
		controller.MeowControllerSet,
		controller.HttpServerSet,
	))
}
