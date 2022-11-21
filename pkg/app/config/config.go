package config

import "rimeks.ru/services/pkg/app/store"

type Config struct {
	Port string        `toml:"port"`
	DB   *store.Config `toml:"db"`
}

func New() *Config {
	return &Config{
		Port: "8080",
		DB:   store.NewPostgresConfig(),
	}
}
