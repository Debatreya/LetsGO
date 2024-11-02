package main

import "fmt"

func main() {
	fmt.Println("Lets Learn About Pointers")

	// Declaring a variable
	var ptr *int // Declaring a pointer variable, `*int` is used to declare a pointer variable that can store the address of an integer variable

	fmt.Println("DEFAULT Value of ptr:", ptr) // Output: DEFAULT Value of ptr: <nil>

	myInt := 2
	ptr = &myInt // Storing the address of `myInt` in `ptr`

	fmt.Println("Address of myInt:", &myInt) // Output: Address of myInt: 0xc0000b6010
	fmt.Println("Value of ptr:", ptr) // Output: Value of ptr: 0xc0000b6010

	// Dereferencing a pointer
	fmt.Println("Value at the address stored in ptr:", *ptr) // Output: Value at the address stored in ptr: 2

	// Changing the value of the variable using the pointer
	*ptr *= 4 // Changing the value of `myInt` using the pointer
	fmt.Println("Value of myInt after changing using ptr:", myInt) // Output: Value of myInt after changing using ptr: 3
}