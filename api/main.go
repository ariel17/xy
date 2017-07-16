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

	s, err := storage.CreateStorage()
	if err != nil {
		log.Fatal("There was an error connecting to database:", err)
		os.Exit(config.ErrorInvalidConfiguration)
	}

	defer s.Close()

	server.Configure()
	log.Fatal(server.Start("0.0.0.0", port))
}
