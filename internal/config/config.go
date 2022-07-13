package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	httpAddress         string
	grpcUsersServiceUrl string
}

func NewAppConfig() (*AppConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &AppConfig{
		httpAddress:         fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")),
		grpcUsersServiceUrl: os.Getenv("GRPC_USERS_SERVICE"),
	}, nil
}

func (c *AppConfig) HttpAddress() string {
	return c.httpAddress
}

func (c *AppConfig) GrpcUsersServiceUrl() string {
	return c.grpcUsersServiceUrl
}
