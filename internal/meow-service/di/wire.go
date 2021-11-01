//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/command"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/config"
	controller "github.com/kshvyryaev/cyber-meower/internal/meow-service/controller/http"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/event"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/repository"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/service"
	"go.uber.org/zap"
)

func InitializeHttpServer(logger *zap.Logger) (*controller.HttpServer, func(), error) {
	panic(wire.Build(
		config.ConfigSet,
		service.MeowTranslatorServiceSet,
		repository.PostgresConnectionSet,
		repository.PostgresMeowRepositorySet,
		event.NatsEventPublisherSet,
		command.СreateMeowCommandHandlerSet,
		controller.MeowControllerSet,
		controller.ErrorHandlerMiddlewareSet,
		controller.RecoveryHandlerMiddlewareSet,
		controller.HttpServerSet,
	))
}
