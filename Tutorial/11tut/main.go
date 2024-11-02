package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in GoLang")

	// Declaring an array
	var fruitList [4]string // Declaring an array of size 4 (Providing size is mandatory)

	// Assigning values to the array
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Grapes"

	// Printing the array
	fmt.Println("Fruit List: ", fruitList) // [Apple Banana  Grapes] -> Notice the empty space at index 2
	fmt.Println("Fruit List Length: ", len(fruitList)) // 4 -> Length of the array
	// What is at index 2? -> Nothing, it is empty
	fmt.Println("Fruit at index 2: ", fruitList[2]) // -> Empty string
	// Are you sure its a string?
	fmt.Printf("Data type of fruitList[2]: %T\n", fruitList[2]) // -> string

	// Declaring and initializing an array
	veggieList := [3]string{"Potato", "Tomato", "Onion"}
	// OR
	var veggieList2 = [3]string{"Potato", "Beans", "Pumpkin"}

	// Printing the array
	fmt.Println("Veggie List: ", veggieList) // [Potato Tomato Onion]
	fmt.Println("Veggie List 2: ", veggieList2) // [Potato Beans Pumpkin]

}