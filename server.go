package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	port = 8080
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/register", formHandler)
	http.HandleFunc("/test", testHandler)

	fmt.Printf("Starting web server at port 8080\n")
	if err := http.ListenAndServe(port); err != nil {
		log.Fatal(err)
	}
}