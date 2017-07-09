package server

import (
	"fmt"
	"log"
	"net/http"
)

const (
	SUBJECTS_PATH string = "/subjects"
	REGISTER_PATH string = "/register"
)

// Configure Maps URL paths into handlers.
func Configure() {
	http.HandleFunc(SUBJECTS_PATH, subjects)
	http.HandleFunc(REGISTER_PATH, register)
}

// Start Servers API endpoint in indicated ip and port.
func Start(ip string, port string) error {
	address := fmt.Sprintf("%s:%s", ip, port)
	log.Println("Starting server in", address)
	return http.ListenAndServe(address, nil)
}
