package main

import (
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/di"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("cannot initialize zap logger: " + err.Error())
	}
	defer logger.Sync()

	server, cleanup, err := di.InitializeHttpServer(logger)
	if err != nil {
		panic("cannot initialize http server: " + err.Error())
	}
	defer cleanup()

	server.Run()
}
