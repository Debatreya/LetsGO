package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App")
	fmt.Println("Please rate our pizza between 1 and 5")

	// Initialize the reader
	reader := bufio.NewReader(os.Stdin)

	// Input Variable
	input, _ := reader.ReadString('\n')

	// Output
	fmt.Println("Thanks for rating, ", input)

	// We want to add 1 to the users input
	// We need to convert the input to an integer. Input is a string now

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // TrimSpace to trim out "4.6" from "4.6\r\n"

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
		// Got an error
		// strconv.ParseFloat: parsing "4.6\r\n": invalid syntax
		// So we have to Trim the input
	}

}