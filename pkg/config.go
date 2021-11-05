package pkg

import (
	"flag"
)

type Config struct {
	Port                     string
	DatabaseConnectionString string
	EventStoreAddress        string
}

func ProvideConfig() *Config {
	config := &Config{
		Port:                     *flag.String("port", "8080", "Server port"),
		DatabaseConnectionString: *flag.String("databaseConnectionString", "host=localhost port=5432 user=postgres password=postgres dbname=cybermeowerdb sslmode=disable", "Database connection string"),
		EventStoreAddress:        *flag.String("eventStoreAddress", "127.0.0.1:4222", "Event store address"),
	}

	flag.Parse()

	return config
}
