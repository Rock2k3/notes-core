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

func NewServer(l logger.AppLogger) *server {
	return &server{
		config:     config.GetAppConfig(),
		logger:     l,
		httpServer: NewHttpServer(),
	}
}

func (s server) Run() error {
	err := s.httpServer.Run(s.config)
	if err != nil {
		return err
	}

	return nil
}
