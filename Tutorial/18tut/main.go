package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions in Golang")

    // Function Call
    greeter()

    // We can not write function inside a function

    // Add two numbers
    res := adder(10, 20)
    fmt.Println("Result:", res)

    // Add multiple numbers
    res, myMessage := proAdder(10, 20, 30, 40, 50)
    fmt.Println("Result:", res)
    fmt.Println("Message:", myMessage)
}

// Function Definition
func greeter() {
    fmt.Println("Namaste! Welcome to Golang Functions")
}

func adder(valOne int, valTwo int) int {
    return valOne + valTwo
}

// ProAdder function with multiple return values
func proAdder(values ...int) (int, string) { // This is a variadic function, `...` is called variadic parameter or ellipsis
    total := 0

    for _, val := range values {
        total += val
    }

    return total, "Success"
}