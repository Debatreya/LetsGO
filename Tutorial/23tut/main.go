package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?course=reactjs&paymentid=bsiuhfsdihfisdf"

func main() {
	fmt.Println("Welcome to Handling URL with GO")
	fmt.Println(myURL)
	
	// Parsing the URL
	result, _ := url.Parse(myURL) // _ is used to ignore the error as I am sure that there will be no error
	// Printing the result
	fmt.Println(result.Scheme) // Scheme is the protocol used in the URL (https)
	fmt.Println(result.Host) // Host is the domain name (lco.dev:3000)
	fmt.Println(result.Path) // Path is the path of the URL (/learn)
	fmt.Println(result.RawQuery) // RawQuery is the raw query parameters in the URL (course=reactjs&paymentid=bsiuhfsdihfisdf)
	fmt.Println(result.Query()) // Query is the query parameters in the URL (map[course:[reactjs] paymentid:[bsiuhfsdihfisdf]])


	// Getting the query parameters
	qParams := result.Query()

	for _, val := range qParams {
		fmt.Println("Param: ", val)
	}

	// parts of the URL
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=abc",
	}

	anotherUrl := partsOfUrl.String() // string(partsOfUrl) will also work
	fmt.Println(anotherUrl)
}