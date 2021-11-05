package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg"
)

type HttpServer struct {
	config                    *pkg.Config
	meowController            *HttpMeowController
	errorHandlerMiddleware    *HttpErrorHandlerMiddleware
	recoveryHandlerMiddleware *HttpRecoveryHandlerMiddleware
}

func ProvideHttpServer(
	config *pkg.Config,
	meowController *HttpMeowController,
	errorHandlerMiddleware *HttpErrorHandlerMiddleware,
	recoveryHandlerMiddleware *HttpRecoveryHandlerMiddleware) *HttpServer {
	return &HttpServer{
		config:                    config,
		meowController:            meowController,
		errorHandlerMiddleware:    errorHandlerMiddleware,
		recoveryHandlerMiddleware: recoveryHandlerMiddleware,
	}
}

func (server *HttpServer) Run() {
	router := gin.New()

	router.Use(server.recoveryHandlerMiddleware.Handle())
	router.Use(server.errorHandlerMiddleware.Handle())

	server.meowController.Route(router)
	router.Run(":" + server.config.Port)
}
