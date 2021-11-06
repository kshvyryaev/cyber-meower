//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/command"
	grpcController "github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/grpc"
	httpController "github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/http"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/event"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/repository"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/service"
)

func InitializeHttpServer() (*httpController.HttpServer, func(), error) {
	panic(wire.Build(
		pkg.ProvideConfig,
		pkg.ProvideZap,
		service.ProvideMeowTranslatorService,
		repository.ProvidePostgres,
		repository.PostgresMeowRepositorySet,
		event.ProvideNats,
		event.NatsMeowEventPublisherSet,
		command.ProvideСreateMeowCommandHandler,
		httpController.ProvideHttpMeowController,
		httpController.ProvideHttpErrorHandlerMiddleware,
		httpController.ProvideHttpRecoveryHandlerMiddleware,
		httpController.ProvideHttpServer,
	))
}

func InitializeGrpcServer() (*grpcController.GrpcServer, func(), error) {
	panic(wire.Build(
		pkg.ProvideConfig,
		pkg.ProvideZap,
		service.ProvideMeowTranslatorService,
		repository.ProvidePostgres,
		repository.PostgresMeowRepositorySet,
		event.ProvideNats,
		event.NatsMeowEventPublisherSet,
		command.ProvideСreateMeowCommandHandler,
		grpcController.ProvideGrpcMeowController,
		grpcController.ProvideGrpcErrorHandlerInterceptor,
		grpcController.ProvideGrpcServer,
	))
}
