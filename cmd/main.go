package main

import (
	"github.com/kshvyryaev/cyber-meower-meower-service/pkg/di"
)

func main() {
	server, cleanup, err := di.InitializeGrpcServer()
	if err != nil {
		panic("cannot initialize http server: " + err.Error())
	}
	defer cleanup()

	err = server.Run()
	if err != nil {
		panic("cannot run http server: " + err.Error())
	}
}
