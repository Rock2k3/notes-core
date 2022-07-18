package server

import (
	"github.com/Rock2k3/notes-core/internal/config"
	"github.com/Rock2k3/notes-core/internal/server/controllers"
	"github.com/Rock2k3/notes-core/pkg/noteshttpserver"
)

type server struct {
	config  *config.AppConfig
	httpSrv *noteshttpserver.HttpServer
}

func NewServer(c *config.AppConfig) *server {
	return &server{config: c, httpSrv: noteshttpserver.NewHttpServer()}
}

func (s *server) Run() error {
	userRoutes := controllers.GetUserRoutes()
	s.httpSrv.AddRoutes(userRoutes)

	return s.httpSrv.Start(s.config.HttpAddress())
}
