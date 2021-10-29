package main

import (
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/di"
)

func main() {
	server, _, err := di.InitializeHttpServer()
	if err != nil {
		// TODO: Добавить логгирование
	}

	server.Run()
}
