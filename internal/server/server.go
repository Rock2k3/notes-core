package server

import (
	"notes-core/internal/config"
	"notes-core/internal/server/controllers"
	"notes-core/pkg/notesHttpServer"
)

type server struct {
	config  *config.AppConfig
	httpSrv *notesHttpServer.HttpServer
}

func NewServer(c *config.AppConfig) *server {
	return &server{config: c, httpSrv: notesHttpServer.NewHttpServer()}
}

func (s *server) Run() error {
	userRoutes := controllers.GetUserRoutes()
	s.httpSrv.AddRoutes(userRoutes)

	return s.httpSrv.Start(s.config.HttpAddress())
}
