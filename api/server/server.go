package server

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// SubjectsPath TODO
	SubjectsPath string = "/subjects"
	// RegisterPath TODO
	RegisterPath string = "/register"
)

// Configure Maps URL paths into handlers.
func Configure() {
	http.HandleFunc(SubjectsPath, Subjects)
	http.HandleFunc(RegisterPath, Register)
}

// Start Servers API endpoint in indicated ip and port.
func Start(ip string, port string) error {
	address := fmt.Sprintf("%s:%s", ip, port)
	log.Println("Starting server in", address)
	return http.ListenAndServe(address, nil)
}
