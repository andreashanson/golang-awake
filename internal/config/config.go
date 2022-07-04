package config

import "os"

type Config struct {
	Server   *ServerConfig
	Postgres *DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

func NewConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
		Postgres: &DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}
}
