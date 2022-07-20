package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddHealthCheckRoutes(router *echo.Router) {
	router.Add(http.MethodGet, "/health_check", handlerHealthCheck())
}

func handlerHealthCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	}
}
