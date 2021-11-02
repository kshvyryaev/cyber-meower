//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/command"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/config"
	controller "github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/http"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/event"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/repository"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/service"
	"go.uber.org/zap"
)

func InitializeHttpServer(logger *zap.Logger) (*controller.HttpServer, func(), error) {
	panic(wire.Build(
		config.ProvideConfig,
		service.ProvideMeowTranslatorService,
		repository.ProvidePostgresDatabase,
		repository.PostgresMeowRepositorySet,
		event.ProvideNatsConnection,
		event.NatsMeowEventPublisherSet,
		command.ProvideСreateMeowCommandHandler,
		controller.ProvideMeowController,
		controller.ProvideErrorHandlerMiddleware,
		controller.ProvideRecoveryHandlerMiddleware,
		controller.ProvideHttpServer,
	))
}
