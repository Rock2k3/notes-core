package config

import (
	"fmt"
	"os"
)

type AppConfig struct {
	logLevel string
	http     HttpConfig
	grpc     GrpcConfig
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (c AppConfig) GetConfig() *AppConfig {
	httpCfg := HttpConfig{
		httpAddress: fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")),
	}
	grpcCfg := GrpcConfig{
		grpcUsersServiceAddress: os.Getenv("GRPC_USERS_SERVICE_ADDRESS"),
	}
	return &AppConfig{
		logLevel: os.Getenv("LOG_LEVEL"),
		http:     httpCfg,
		grpc:     grpcCfg,
	}
}

func (c *AppConfig) LogLevel() string {
	return c.logLevel
}

func (c *AppConfig) Http() HttpConfig {
	return c.http
}

func (c *AppConfig) Grpc() GrpcConfig {
	return c.grpc
}
