package routes

import (
	"github.com/Rock2k3/notes-core/internal/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterHealthCheckRoutes(router *echo.Router) {
	log := logger.GetAppLogger()
	log.Debug("RegisterHealthCheckRoutes")

	router.Add(http.MethodGet, "/health_check", handlerHealthCheck())
}

func handlerHealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	}
}
