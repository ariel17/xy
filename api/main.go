package main

import (
	"./server"
	"log"
	"os"
)

func main() {
	port := os.Args[1]

	server.Configure()
	log.Fatal(server.Start("0.0.0.0", port))
}
