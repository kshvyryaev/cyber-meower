package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg"
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/controller/http/middleware"
)

type HttpServer struct {
	config                    *pkg.Config
	meowController            *HttpMeowController
	errorHandlerMiddleware    *middleware.HttpErrorHandlerMiddleware
	recoveryHandlerMiddleware *middleware.HttpRecoveryHandlerMiddleware
}

func ProvideHttpServer(
	config *pkg.Config,
	meowController *HttpMeowController,
	errorHandlerMiddleware *middleware.HttpErrorHandlerMiddleware,
	recoveryHandlerMiddleware *middleware.HttpRecoveryHandlerMiddleware) *HttpServer {
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
