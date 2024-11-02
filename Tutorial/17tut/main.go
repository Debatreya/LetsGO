package main

import "fmt"

func main() {
	fmt.Println("Learn about Loops in Golang")

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For loop with Slice
	Days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for i := 0; i < len(Days); i++ {
		fmt.Println(Days[i])
	}
	// For using range keyword
	for idx, day := range Days { // idx is index and day is value // We can ignore index by using _ (underscore)
		fmt.Println(idx, ":", day)
	}
	// OR
	for i := range Days { // i is index
		fmt.Println(Days[i])
	}

	// Goto using label
	val := 1

	for val <= 10 {
		if val == 8 {
			goto lco // Goto statement
		}
		if val == 5 {
			fmt.Println("Value is 5")
			val++ // Incrementing value
		}
		fmt.Println(val)
		val++ // Incrementing value // ++val is not allowed in Golang
	}

lco:
	fmt.Println("Hello LCO")

}
