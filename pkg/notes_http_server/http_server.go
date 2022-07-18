package notes_http_server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

type HttpServer struct {
	*echo.Echo
}

func NewHttpServer() *HttpServer {
	s := echo.New()
	s.GET("/health_check", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})

	return &HttpServer{s}
}

func (s HttpServer) AddRoutes(r []Route) {
	for _, r := range r {
		s.Add(r.Method, r.Path, r.Handler)
	}
}
