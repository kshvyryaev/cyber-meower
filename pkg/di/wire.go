// go:build wireinject
//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/grpc"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/http"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/http/middleware"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/domain"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/event"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/repository"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/service"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/usecase"
)

func InitializeHttpServer() (*http.HttpServer, func(), error) {
	panic(wire.Build(
		pkg.ProvideConfig,
		pkg.ProvideZap,
		service.ProvideMeowTranslatorService,
		wire.Bind(new(domain.MeowTranslatorService), new(*service.MeowTranslatorService)),
		repository.ProvidePostgres,
		repository.ProvidePostgresMeowRepository,
		wire.Bind(new(domain.MeowRepository), new(*repository.PostgresMeowRepository)),
		event.ProvideNats,
		event.ProvideNatsMeowEventPublisher,
		wire.Bind(new(domain.MeowEventPublisher), new(*event.NatsMeowEventPublisher)),
		usecase.ProvideMeowUsecase,
		wire.Bind(new(domain.MeowUsecase), new(*usecase.MeowUsecase)),
		middleware.ProvideHttpErrorHandlerMiddleware,
		middleware.ProvideHttpRecoveryHandlerMiddleware,
		http.ProvideHttpMeowController,
		http.ProvideHttpServer,
	))
}

func InitializeGrpcServer() (*grpc.GrpcServer, func(), error) {
	panic(wire.Build(
		pkg.ProvideConfig,
		pkg.ProvideZap,
		service.ProvideMeowTranslatorService,
		wire.Bind(new(domain.MeowTranslatorService), new(*service.MeowTranslatorService)),
		repository.ProvidePostgres,
		repository.ProvidePostgresMeowRepository,
		wire.Bind(new(domain.MeowRepository), new(*repository.PostgresMeowRepository)),
		event.ProvideNats,
		event.ProvideNatsMeowEventPublisher,
		wire.Bind(new(domain.MeowEventPublisher), new(*event.NatsMeowEventPublisher)),
		usecase.ProvideMeowUsecase,
		wire.Bind(new(domain.MeowUsecase), new(*usecase.MeowUsecase)),
		grpc.ProvideGrpcMeowController,
		grpc.ProvideGrpcErrorHandlerInterceptor,
		grpc.ProvideGrpcServer,
	))
}
