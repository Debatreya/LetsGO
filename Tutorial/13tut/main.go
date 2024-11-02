package main

import (
	"fmt"
)

func main() {
	fmt.Println("Lets Learn About MAPs")

	// Creating a map
	// Using make function
	// Syntax: make(map[key-type]value-type)

	languages := make(map[string]string)

	// Adding key-value pairs
	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"
	languages["go"] = "Golang"

	fmt.Println("Languages:", languages)
	fmt.Println("js shorts for", languages["js"])

	// Deleting a key-value pair
	delete(languages, "rb")

	fmt.Println("After deleing 'rb' key Languages:", languages)

	// Looping in Maps
	for key, value := range languages {
		fmt.Println(key, ":", value)
	}

	// If we don't care about the key (, OK syntax)
	for _, value := range languages {
		fmt.Println(value)
	}
}