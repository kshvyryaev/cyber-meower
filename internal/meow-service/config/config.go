package config

import "flag"

type Config struct {
	Database *DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func NewConfig() *Config {
	return &Config{
		Database: &DatabaseConfig{
			Host:     *flag.String("database-host", "localhost", "Database host"),
			Port:     *flag.String("database-port", "5432", "Database port"),
			User:     *flag.String("database-user", "postgres", "Database user"),
			Password: *flag.String("database-password", "postgres", "Database password"),
			Name:     *flag.String("database-name", "cybermeowerdb", "Database name"),
			SSLMode:  *flag.String("database-sslmode", "disable", "Database SSL mode"),
		},
	}
}
