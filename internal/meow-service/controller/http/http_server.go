package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/config"
)

type HttpServer struct {
	Config         *config.Config
	MeowController *MeowController
}

func (server *HttpServer) Run() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(RecoveryHandlerMiddleware())
	router.Use(ErrorHandlerMiddleware())

	server.MeowController.Route(router)
	router.Run(":" + server.Config.Port)
}

func ProvideHttpServer(config *config.Config, meowController *MeowController) *HttpServer {
	return &HttpServer{
		Config:         config,
		MeowController: meowController,
	}
}

var HttpServerSet = wire.NewSet(ProvideHttpServer)
