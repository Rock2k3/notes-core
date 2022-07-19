package server

import (
	"github.com/Rock2k3/notes-core/internal/appV2/config"
	"github.com/labstack/echo/v4"
)

type httpServer struct {
	echo *echo.Echo
}

func NewHttpServer() *httpServer {
	return &httpServer{echo: echo.New()}
}

func (s httpServer) Run(c *config.AppConfig) error {
	return s.echo.Start(c.Http().HttpAddress())
}
