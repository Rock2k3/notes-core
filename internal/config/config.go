package config

import (
	"fmt"
	"os"
)

var config *AppConfig

type AppConfig struct {
	logLevel string
	http     httpConfig
	grpc     grpcConfig
}

func GetAppConfig() *AppConfig {
	return config
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (c AppConfig) Init() {
	httpCfg := httpConfig{
		httpAddress: fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")),
	}
	grpcCfg := grpcConfig{
		grpcUsersServiceAddress: os.Getenv("GRPC_USERS_SERVICE_ADDRESS"),
	}

	config = &AppConfig{
		logLevel: os.Getenv("LOG_LEVEL"),
		http:     httpCfg,
		grpc:     grpcCfg,
	}
}

func (c *AppConfig) LogLevel() string {
	return c.logLevel
}

func (c *AppConfig) Http() *httpConfig {
	return &c.http
}

func (c *AppConfig) Grpc() *grpcConfig {
	return &c.grpc
}
