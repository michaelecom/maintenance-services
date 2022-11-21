package main

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"rimeks.ru/services/pkg/app/config"
	"rimeks.ru/services/pkg/app/handler"
	"rimeks.ru/services/pkg/app/server"
	"rimeks.ru/services/pkg/app/service"
	"rimeks.ru/services/pkg/app/store"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.toml", "path to config file")
}

func main() {
	flag.Parse()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	config := config.New()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	db, err := store.NewPostgresDB(store.Config{
		Host:     config.DB.Host,
		Port:     config.DB.Port,
		DBName:   config.DB.DBName,
		SSLMode:  config.DB.SSLMode,
		Username: config.DB.Username,
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})
	if err != nil {
		logrus.Fatal(err)
	}

	store := store.New(db)
	service := service.New(store)
	handler := handler.New(service)

	server := new(server.Server)

	if err := server.Run(config.Port, handler.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}
