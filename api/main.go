package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ariel17/xy/api/controllers"
)

func main() {
	http.HandleFunc("/users", controllers.Users)
	http.HandleFunc("/register", controllers.Register)

	address := fmt.Sprintf("0.0.0.0:%s", os.Args[1])
	log.Println("Starting server in", address)

	log.Fatal(http.ListenAndServe(address, nil))
}
