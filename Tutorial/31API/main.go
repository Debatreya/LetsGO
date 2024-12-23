package main

import (
	"fmt"
	"goAPImod/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API - MongoDB - Netflix")
	r := router.Router()
	fmt.Println("Server is getting ready...")
	fmt.Println("Server is running on port 4000 | http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}