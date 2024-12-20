package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Making and Handling GET request")
	// Declare test URL
	const URL = "https://jsonplaceholder.typicode.com/posts/1"
	PerformGetRequest(URL)
}

func PerformGetRequest(URL string) {
	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Response Status: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength) // -1 if unknown

	content, _ := io.ReadAll(response.Body) // content is a byte array
	// fmt.Println(string(content))

	// Use Builder to build string
	var responseString strings.Builder // Prefered in handling web content (It gives much control over the content)
	byteCount, _ := responseString.Write(content) // byteCount is the number of bytes written, returned by Write

	fmt.Println("ByteCount: ", byteCount)
	fmt.Println("Response String: ", responseString.String())
}