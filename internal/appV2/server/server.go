package server

import (
	"github.com/Rock2k3/notes-core/internal/appV2/config"
	"go.uber.org/zap"
)

type server struct {
	config *config.AppConfig
}

func NewServer(c *config.AppConfig, logger *zap.Logger) *server {
	logger.Info("NewServer")
	return &server{
		config: c,
	}
}
