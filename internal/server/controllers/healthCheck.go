package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterHealthCheckHandlers(s *echo.Echo) *echo.Echo {
	s.GET("health_check", healthCheck)
	return s
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Ok")
}
