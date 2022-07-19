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
	defer appLogger.Sync()

	s := server.NewServer(appConfig, appLogger)
	err := s.Run()
	if s.Run() != nil {
		appLogger.Fatalf("Error starting server %v", err)
	}

}
