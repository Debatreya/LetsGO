package main

import "fmt"

const LoginToken string = "asdjfsvjhvu3827rfv" // constant variable

// Go is a statically typed language
// This means that we have to specify the type of a variable when we create it
// Go has a lot of built-in data types
func main() {
	// To create a variable we use `var` <variable-name> <data-type> = <value>
	var username string = "Debatreya"
	fmt.Println("Username:", username);
	// To get the type of a variable we use `%T` in `Printf` function like in C
	fmt.Printf("Type of username: %T\n", username);

	// uintN (unsigned integer) is a type of integer which can only store positive values (0 to 2^N). -> uint8, uint16, uint32, uint64
	var age uint8 = 20
	fmt.Println("Age:", age);
	fmt.Printf("Type of age: %T\n", age);

	// Floats are used to store decimal values. -> float32 (5 decimal places), float64 (15 decimal places)
	var smallFloat float32 = 255.9897789707078077070790707
	fmt.Println("Pi:", smallFloat);
	fmt.Printf("Type of smallFloat: %T\n", smallFloat);


	// DEFAULT VALUES and ALIASes
	var anotherInt int
	fmt.Println("Default value of int:", anotherInt); // Default value of int is 0

	// implicit type declaration
	// Go can automatically detect the type of a variable based on the value assigned to it
	// This is called implicit type declaration
	// Syntax: <variable-name> := <value>  // `:=` is called the walrus operator
	// OR
	// Syntax: <variable-name> = <value>
	// The first syntax is preferred
	// The second syntax is used when we want to change the value of a variable

	var website = "google.co.in"
	fmt.Printf("Type of website: %T\n", website); // Type of website: string

	// no var style
	// We can also create a variable without using the `var` keyword
	// This is called the no var style
	// Syntax: <variable-name> := <value>
	// This is only used inside functions
	// We cannot use this outside functions
	// We cannot use this to create global variables
	// We cannot use this to create variables without assigning a value
	// We cannot use this to create variables with a type different from the value assigned to it
	// We cannot use this to create variables with the same name as a global variable
	// We cannot use this to create variables with the same name as another variable in the same function

	// This is valid
	id := "ID456"
	fmt.Println("Id:", id);

	// Accessing global variables
	fmt.Println("Login Token:", LoginToken);
}