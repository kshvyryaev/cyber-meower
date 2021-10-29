package config

import "flag"

type Config struct {
	DatabaseConnectionString string
}

func NewConfig() *Config {
	return &Config{
		DatabaseConnectionString: *flag.String("database-connection-string", "host=localhost port=5432 user=postgres password=postgres dbname=cybermeowerdb sslmode=disable", "Database connection string"),
	}
}
