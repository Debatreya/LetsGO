package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in Golang: Rolling the Dice...")

	// Note: We dont need `break`, and we have `fallthrough`

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	diceNumber := rand.Intn(6) + 1   // 0 - 5, +1 -> 1 - 6

	fmt.Println("Value of dice is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice number is 1 and you can open")
	case 2:
		fmt.Println("Dice number is 2 and you can move 2 steps")
	case 3:
		fmt.Println("Dice number is 3 and you can move 3 steps")
		fallthrough // Execute the next case also
	case 4:
		fmt.Println("Dice number is 4 and you can move 4 steps")
	case 5:
		fmt.Println("Dice number is 5 and you can move 5 steps")
		fallthrough // Execute the next case also
	case 6:
		fmt.Println("Dice number is 6 and you can move 6 steps and roll the dice again")
		main() // Recursive function call
	default:
		fmt.Println("Invalid dice number")
	}
}
