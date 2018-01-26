package main

import (
	"log"
	"os"

	"github.com/ariel17/xy/api/config"
	"github.com/ariel17/xy/api/server"
	"github.com/ariel17/xy/api/storage"
)

func main() {
	port := os.Args[1]

	if err := storage.CreateStorage(); err != nil {
		log.Fatal("error connecting to database", err)
		os.Exit(config.ErrorInvalidConfiguration)
	}
	defer storage.Instance.Close()

	server.Configure()
	log.Fatal(server.Start("0.0.0.0", port))
}
