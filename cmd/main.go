package main

import (
	"github.com/Rock2k3/notes-core/internal/config"
	"github.com/Rock2k3/notes-core/internal/server"
	"log"
)

func main() {
	appConf, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = server.NewServer(appConf).Run()
}
