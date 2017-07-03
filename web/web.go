package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "file", "index.html", "template file")
	flag.Parse()
}

func render(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		fmt.Println("Template parse error: ", err)
		return
	}

	err = t.Execute(w, "")
	if err != nil {
		fmt.Println("template executing error: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", render)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Args[1]), nil))
}
