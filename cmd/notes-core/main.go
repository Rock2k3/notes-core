package main

import (
	"github.com/Rock2k3/notes-core/internal/config"
	"github.com/Rock2k3/notes-core/internal/infra/server"
	"github.com/Rock2k3/notes-core/internal/logger"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.NewAppConfig().Init()
}

func main() {
	appConfig := config.GetAppConfig()

	appLogger := logger.NewAppLogger(appConfig)
	appLogger.Init()
	defer appLogger.Sync()

	s := server.NewServer(appLogger)
	err := s.Run()
	if err != nil {
		appLogger.Fatalf("Error starting server %v", err)
	}

}
