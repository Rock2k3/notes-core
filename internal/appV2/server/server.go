package server

import (
	"github.com/Rock2k3/notes-core/internal/appV2/config"
	"github.com/Rock2k3/notes-core/internal/appV2/logger"
)

type server struct {
	config     *config.AppConfig
	logger     logger.AppLogger
	httpServer *httpServer
}

func NewServer(c *config.AppConfig, l logger.AppLogger) *server {
	return &server{
		config:     c,
		logger:     l,
		httpServer: NewHttpServer(),
	}
}

func (s server) Run() error {
	s.logger.Info("http server starting")
	err := s.httpServer.Run(s.config)
	if err != nil {
		return err
	}

	return nil
}
