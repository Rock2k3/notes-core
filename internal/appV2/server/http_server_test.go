package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpServer_handlerHealthCheck(t *testing.T) {
	testHttpServer := NewHttpServer()
	testHttpServer.configureRouter()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/health_check", nil)
	testHttpServer.echo.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
