package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
	httpAddress             string
	grpcUsersServiceAddress string
	logLevel                string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (c AppConfig) GetConfig() *AppConfig {
	return &AppConfig{
		httpAddress:             fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")),
		grpcUsersServiceAddress: os.Getenv("GRPC_USERS_SERVICE_ADDRESS"),
		logLevel:                os.Getenv("LOG_LEVEL"),
	}
}

func (c *AppConfig) HttpAddress() string {
	return c.httpAddress
}

func (c *AppConfig) GrpcUsersServiceUrl() string {
	return c.grpcUsersServiceAddress
}

func (c *AppConfig) LogLevel() string {
	return c.logLevel
}
