package main

import (
	"github.com/Rock2k3/notes-core/internal/appV2/config"
	"github.com/Rock2k3/notes-core/internal/appV2/logger"
	"github.com/Rock2k3/notes-core/internal/appV2/server"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	appConfig := config.NewAppConfig().GetConfig()

	appLogger := logger.NewAppLogger(appConfig)
	appLogger.Init()
	l := appLogger.Logger()
	defer l.Sync()

	server.NewServer(appConfig, l)
	l.Info("app started")

}
