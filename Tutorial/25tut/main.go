package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Print("Lets Make a POST request with all types of payload (1. JSON, 2. URL Encoded, 3. Form Data, 4. XML, 5. File)")
	// Declare test URL for JSON (Others in the next tutorials)
	const URL = "https://jsonplaceholder.typicode.com/posts"
	JSONbody := strings.NewReader(`
		{
			"title": "foo",
			"body": "bar",
			"userId": 1,
			"id": 101
		}
	`)
	PerformPostRequest(URL, JSONbody)
}

func PerformPostRequest(URL string, body io.Reader) {
	fmt.Printf("\n\nPerforming POST request to URL: %s\n", URL)
	fmt.Printf("Body: %s\n", body)
	// Perform POST request
	// Code to perform POST request
	response, err := http.Post(URL, "application/json", body) // URL, Content-Type, Body
	// Request and Response are always io.Reader type
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body) // content is a byte array
	// Create a builder to store response body
	var postResBuilder strings.Builder
	// Read response body
	_, error := postResBuilder.Write(content)
	if error != nil {
		panic(error)
	}
	fmt.Println("Response: ", postResBuilder.String())
}