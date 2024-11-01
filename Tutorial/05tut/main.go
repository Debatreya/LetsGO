package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input Tutorial"
	fmt.Println(welcome)

	// Declare a buffer reader object
	reader := bufio.NewReader(os.Stdin) // Create a reader object using bufio package NewReader function and os package Stdin variable which says to take input from OS's Input Output
	fmt.Println("Enter your name: ")
	
	// comma ok  || err ok syntax
	input,_ := reader.ReadString('\n') // ReadString function reads the input from the user until it finds the delimiter '\n' which is a newline character
	// Here we are using _ to ignore the error value returned by ReadString function, we can also use comma ok syntax to get the error value
	// input, err := reader.ReadString('\n')
	// Sometimes we dont care about input value and only want to check if there is an error or not, in that case we can use comma ok syntax
	// _, err := reader.ReadString('\n')
	
	fmt.Println("Hello, ", input)
}
