package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	port = ":8080"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/test" {
		http.Error(w, "Not Found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Test Complete!")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.Error(w, "Not Found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported.", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful.")

	username := r.FormValue("email")
	password := r.FormValue("username")
	email := r.FormValue("password")
	fmt.Fprintf(w, "email: %v\n", email)
	fmt.Fprintf(w, "name: %v\n", username)
	fmt.Fprintf(w, "password: %v\n", password)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/test", testHandler)

	fmt.Printf("Starting web server at port 8080\n")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}