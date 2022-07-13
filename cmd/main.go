package main

import (
	"log"
	"notes-core/internal/config"
	"notes-core/internal/server"
)

func main() {
	appConf, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = server.NewServer(appConf).Run()
}
