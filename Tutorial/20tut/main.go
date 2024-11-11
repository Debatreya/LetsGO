package main

import "fmt"

func main() {
	defer fmt.Println("This will be printed at the end")
	// defer is used to delay the execution of a function until the surrounding function returns
	// It is used to delay the execution of a function until the surrounding function returns
	fmt.Println("Lets learn about Defer")
	fmt.Println("This will be printed first")
	deferFn()

	// Defer statements are executed in LIFO order
	counting()
}

// deferFn is a function that has some defer statements
func deferFn() {
	defer fmt.Println("This will be printed at the end of deferFn")
	fmt.Println("This will be printed first in deferFn")
}

func counting(){
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
	fmt.Println("Counting done")
}