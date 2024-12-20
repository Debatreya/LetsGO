package main

import (
	"fmt"
	"io"
	"net/url"
	"strings"
	"net/http"
)

func main() {
	fmt.Println("Lets Make a POST request with FORM Data payload")
	// Declare test URL
	const URL = "https://jsonplaceholder.typicode.com/posts"
	// Declare test form data
	// How to create form data in Go? -> Use url.Values
	data := url.Values{}
	data.Add("title", "foo")
	data.Add("body", "bar")
	data.Add("userId", "1")
	// Perform POST request
	PerformPostRequest(URL, data) 
}

func PerformPostRequest(URL string, body url.Values){
	// Create a new POST request
	// res, err := http.NewRequest("POST", URL, strings.NewReader(body.Encode())) // Always request is type io.Reader
	// // Here data is the form data and we are converting it to io.Reader
	// // But first Encode the form data, i.e. convert it to URL encoded format
	// if err != nil {
	// 	panic(err)
	// }
	
	// There is another way for form data in POST

	response, err := http.PostForm(URL, body) // Here we directly pass the type url.Values as form data
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// Read the response body
	content, _ := io.ReadAll(response.Body)
	// Create a builder to store the response
	var builder strings.Builder
	// Write the response to the builder
	builder.Write(content) // Returns bytecount and error
	// Print the response
	fmt.Println(builder.String())
}