package server

import (
	"github.com/Rock2k3/notes-core/internal/config"
	"github.com/Rock2k3/notes-core/internal/server/controllers"
	"github.com/Rock2k3/notes-core/pkg/notes_http_server"
)

type server struct {
	config  *config.AppConfig
	httpSrv *notes_http_server.HttpServer
}

func NewServer(c *config.AppConfig) *server {
	return &server{config: c, httpSrv: notes_http_server.NewHttpServer()}
}

func (s *server) Run() error {
	userRoutes := controllers.GetUserRoutes()
	s.httpSrv.AddRoutes(userRoutes)

	return s.httpSrv.Start(s.config.HttpAddress())
}
