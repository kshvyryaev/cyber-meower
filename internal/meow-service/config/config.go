package config

import (
	"flag"

	"github.com/google/wire"
)

type Config struct {
	Port                     string
	DatabaseConnectionString string
}

func ProvideConfig() *Config {
	return &Config{
		Port:                     *flag.String("port", "8080", "Server port"),
		DatabaseConnectionString: *flag.String("databaseConnectionString", "host=localhost port=5432 user=postgres password=postgres dbname=cybermeowerdb sslmode=disable", "Database connection string"),
	}
}

var ConfigSet = wire.NewSet(
	ProvideConfig,
)
