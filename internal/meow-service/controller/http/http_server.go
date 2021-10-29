package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type HttpServer struct {
	MeowController *MeowController
}

func (server *HttpServer) Run() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(RecoveryHandlerMiddleware())
	router.Use(ErrorHandlerMiddleware())

	server.MeowController.Route(router)
	router.Run()
}

func ProvideHttpServer(meowController *MeowController) *HttpServer {
	return &HttpServer{MeowController: meowController}
}

var HttpServerSet = wire.NewSet(ProvideHttpServer)
