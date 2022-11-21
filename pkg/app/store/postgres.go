package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	ongoingMaintenanceTable = "ongoing_maintenance"
)

type Config struct {
	Host     string
	Port     string
	DBName   string
	Username string
	Password string
	SSLMode  string
}

func NewPostgresConfig() *Config {
	return &Config{}
}

func NewPostgresDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		config.Host, config.Port, config.DBName, config.Username, config.Password, config.SSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
