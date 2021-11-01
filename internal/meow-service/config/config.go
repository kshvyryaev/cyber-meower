package config

import (
	"flag"

	"github.com/google/wire"
)

type Config struct {
	Port                     string
	DatabaseConnectionString string
	EventStoreAddress        string
}

func ProvideConfig() *Config {
	return &Config{
		Port:                     *flag.String("port", "8080", "Server port"),
		DatabaseConnectionString: *flag.String("databaseConnectionString", "host=localhost port=5432 user=postgres password=postgres dbname=cybermeowerdb sslmode=disable", "Database connection string"),
		EventStoreAddress:        *flag.String("eventStoreAddress", "127.0.0.1:4222", "Event store address"),
	}
}

var ConfigSet = wire.NewSet(
	ProvideConfig,
)
