package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ariel17/xy/api/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/users/:id", controllers.GetUsers)
	router.POST("/users", controllers.PostUsers)

	log.Fatal(http.ListenAndServe(":8080", router))

	address := fmt.Sprintf("0.0.0.0:%s", os.Args[1])
	log.Println("Starting server in", address)

	log.Fatal(http.ListenAndServe(address, router))
}
