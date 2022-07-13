package server

import (
	"github.com/labstack/echo/v4"
	"notes-core/internal/config"
	"notes-core/internal/server/controllers"
)

type server struct {
	config  *config.AppConfig
	httpSrv *echo.Echo
}

func NewServer(c *config.AppConfig) *server {
	return &server{config: c, httpSrv: echo.New()}
}

func (s *server) Run() error {
	controllers.RegisterHealthCheckHandlers(s.httpSrv)
	controllers.RegisterUserHandlers(s.httpSrv)

	return s.httpSrv.Start(s.config.HttpAddress())
}
