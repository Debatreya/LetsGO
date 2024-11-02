package main

import "fmt"

func main() {
	fmt.Println("If Else in GoLang")

	// If Else in GoLang
	loginCount := 10
	var result string

	if loginCount < 10 { // We can not move `{` to next line
		result = "Regular User"
	} else if loginCount > 10 { // We can not move `else if` to next line
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	// Another way to write if else

	if 9%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}


	// Short statement
	if num := 3; num < 10 { // num is only available in this block of code and it's scope is limited to this block
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

	// fmt.Println(num) // This will give error as num is not available outside the block
}
