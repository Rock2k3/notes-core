package server

import (
	"github.com/Rock2k3/notes-core/internal/appV2/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

type httpServer struct {
	echo   *echo.Echo
	router *echo.Router
}

func NewHttpServer() *httpServer {
	e := echo.New()
	return &httpServer{
		echo:   e,
		router: e.Router(),
	}
}

func (s httpServer) Run(c *config.AppConfig) error {
	s.configureRouter()
	return s.echo.Start(c.Http().HttpAddress())
}

func (s httpServer) configureRouter() {
	s.router.Add(http.MethodGet, "/health_check", s.HealthCheckHandler())
}

func (s httpServer) HealthCheckHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	}
}
