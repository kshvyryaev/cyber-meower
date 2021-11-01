package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/config"
)

type HttpServer struct {
	Config                    *config.Config
	MeowController            *MeowController
	ErrorHandlerMiddleware    *ErrorHandlerMiddleware
	RecoveryHandlerMiddleware *RecoveryHandlerMiddleware
}

func (server *HttpServer) Run() {
	router := gin.New()

	router.Use(server.RecoveryHandlerMiddleware.Handle())
	router.Use(server.ErrorHandlerMiddleware.Handle())

	server.MeowController.Route(router)
	router.Run(":" + server.Config.Port)
}

func ProvideHttpServer(
	config *config.Config,
	meowController *MeowController,
	errorHandlerMiddleware *ErrorHandlerMiddleware,
	recoveryHandlerMiddleware *RecoveryHandlerMiddleware) *HttpServer {
	return &HttpServer{
		Config:                    config,
		MeowController:            meowController,
		ErrorHandlerMiddleware:    errorHandlerMiddleware,
		RecoveryHandlerMiddleware: recoveryHandlerMiddleware,
	}
}

var HttpServerSet = wire.NewSet(ProvideHttpServer)
