package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Lets Learn about Slice in Golang")

	// Slice is a reference to an array
	// Slice is a lightweight data structure that gives a view of an array
	// Slice is a variable length array
	// Slice is a segment of an array
	// Slice is a dynamic array

	// Declaring a slice
	var fruitList = []string{} // Empty slice `{}` is extra, different also, size is not mentioned.
	fmt.Println("Fruit List is: ", fruitList)
	// Type of fruitList is []string
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// Inserting elements in slice
	// We cannot directly use index to insert elements in slice
	// Use append() function to insert elements in slice
	fruitList = append(fruitList, "Apple")
	fruitList = append(fruitList, "Banana", "Mango")

	fmt.Println("Updated Fruit List is: ", fruitList)

	// Slice of Slice
	fruitList = append(fruitList[1:]) // Removes first element from slice -> Banana, Mango

	fmt.Println("Updated Fruit List after slicing is: ", fruitList)

	fruitList = append(fruitList[:len(fruitList)-1]) // Removes last element from slice -> Banana

	fmt.Println("Updated Fruit List after slicing is: ", fruitList)


	// Creating a slice using make() function
	// make() function is used to create slices, maps, and channels
	// make() function returns a reference to the zero value of the type
	// make() function is used to create slices with non-zero length
	// make() function is used to create slices with non-zero capacity
	// make() function is used to create slices with non-zero length and capacity
	// make() function is used to create slices with non-zero length and capacity

	// Syntax: make([]T, length, capacity)
	// T: Type of slice
	// length: number of elements in slice
	// capacity: number of elements in underlying array of slice

	// Capacity vs Length
	// Length: number of elements in slice
	// Capacity: number of elements in underlying array of slice
	// Length <= Capacity

	// DEMONSTRATION
	// Creating highScores slice using make() function with length 4 and capacity 10
	highScores := make([]int, 4, 10)
	fmt.Println("High Scores: ", highScores)
	fmt.Printf("Length of highScores: %d\n", len(highScores))
	fmt.Printf("Capacity of highScores: %d\n", cap(highScores))

	// Appending elements in highScores slice
	highScores[0] = 234
	highScores[1] = 945
	highScores[3] = 867

	fmt.Println("High Scores: ", highScores)

	// Appending more elements in highScores slice
	highScores = append(highScores, 555, 678, 123, 456, 789)
	fmt.Println("High Scores: ", highScores)
	fmt.Printf("Length of highScores: %d\n", len(highScores))
	fmt.Printf("Capacity of highScores: %d\n", cap(highScores))

	// Append More elements
	highScores = append(highScores, 459, 100)
	fmt.Println("High Scores: ", highScores)
	fmt.Printf("Length of highScores: %d\n", len(highScores)) // 11 where as Initially Defined Capacity was 10
	fmt.Printf("Capacity of highScores: %d\n", cap(highScores)) // as len > cap -> cap = cap*2 => 20

	// WHAT IF WE DONT MENTION CAPACITY??
	lowScores := make([]int, 4)
	// Remember index are used to Overwrite if index already exists.
	// Hence lowScores[4] = 10 -> ERROR
	// But append() can extend the length and put new values.

	fmt.Println("Low Scores: ", lowScores)
	fmt.Printf("Length of lowScores: %d\n", len(lowScores))
	fmt.Printf("Capacity of lowScores: %d\n", cap(lowScores)) // Intially capacity = length = 4

	lowScores[3] = 7
	fmt.Println("Updated by index : Low Scores: ", lowScores)

	// Try append()
	lowScores = append(lowScores, 1, 3)
	fmt.Println("Updated using append(): Low Scores: ", lowScores)
	fmt.Printf("Length of lowScores: %d\n", len(lowScores)) // 6
	fmt.Printf("Capacity of lowScores: %d\n", cap(lowScores)) // as 6 > 4(initially) -> 4*2 => 8
	
	// SORTING
	if !sort.IntsAreSorted(lowScores) {
		fmt.Println("Seems like lowScores is not Sorted Lets Sort It")
		sort.Ints(lowScores)
	}

	fmt.Println("Sorted Low Scores", lowScores)
	
	// REMOVING VALUES FROM SLICES

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("List of Courses", courses)
	// Deleting the element at index
	var index int = 2
	courses = append(courses[:index], courses[index+1:]... ) // ... is spread operator or eclipses
	fmt.Println("Updated Course List", courses)
}
