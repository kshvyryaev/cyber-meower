package main

import (
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/config"
	"github.com/kshvyryaev/cyber-meower/internal/meow-service/di"
)

func main() {
	config := config.NewConfig()

	server, _, err := di.InitializeHttpServer(config)
	if err != nil {
		// TODO: Добавить обработку ошибок
	}

	server.Run()
}
