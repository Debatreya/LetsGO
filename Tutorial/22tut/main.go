package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
)

const url = "https://www.piyushgarg.dev"

func main() {
	fmt.Println("Welcome to Web Request Handling Tutorial in Golang")

	// Making a GET request
	response, err := http.Get(url)

	checkError(err)
	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // Caller should close the response body when done

	databytes, err := io.ReadAll(response.Body)
	checkError(err)

	content := string(databytes)
	// fmt.Println(content)
	// Write the content to a file
	file, err := os.Create("./piyushgarg.html")
	checkError(err)
	defer file.Close()
	_, err = io.WriteString(file, content)
	checkError(err)
	fmt.Println("File Written Successfully: ", file)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error Occurred: ", err)
	}
}