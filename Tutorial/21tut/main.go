package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in GoLang")

	content := "This needs to go in a file - Debatreya Das"

	// Writing to a file
	file, err := os.Create("./myFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNil(err)

	length, err := io.WriteString(file, content)
	checkNil(err)

	fmt.Println("Length is: ", length)
	defer file.Close() // Defer is used to close the file after the function completes

	// Reading from a file
	readFile("./myFile.txt")

}

func readFile(filename string){
	// databyte, err := ioutil.ReadFile(filename) //Read in Bytes format and return the bytes
	// ioutil.ReadFile() is deprecated, use os.ReadFile() instead

	databyte, err := os.ReadFile(filename) //Read in Bytes format and return the bytes

	checkNil(err)

	fmt.Println("Text data inside the file is: ", string(databyte)) // Convert the bytes to string
} 


func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}