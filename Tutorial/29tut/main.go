package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	// Gorilla mux
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// http.ListenAndServe(":4098", r) // Pass port number and router
	log.Fatal(http.ListenAndServe(":4098", r))
}

func greeter() {
	fmt.Println("Hello greeter")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go Web Development</h1>"))
}