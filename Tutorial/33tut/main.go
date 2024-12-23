package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // This is a WaitGroup object (from package sync) // pointer

// Declaring a mutex
var mut sync.Mutex // pointer

func main() {
	fmt.Println("Lets learn about Parallelism in Go -> Goroutines")
	// goroutines are lightweight threads of execution
	// goroutines are functions that run concurrently with other functions
	// goroutines are created using the go keyword followed by a function invocation
	// goroutines are multiplexed into a number of OS threads that run on a smaller number of OS threads
	// goroutines are managed by the Go runtime
	// goroutines are used to run functions concurrently with other functions in the same address space 
	// To use goroutines, we use the go keyword followed by a function invocation

	// Example 1: This is Sequential Execution
	// greeter("Hello")
	// greeter("World")

	// Example 2 (check output yourself) -> No "Hello" is Printed : Why? -> Because the main function exits before the goroutine is executed
	// go greeter("Hello")
	// greeter("World")

	// Example 3: This is Parallel Execution
	websitelist := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.flipkart.com",
		"https://www.instagram.com",
		"https://www.linkedin.com",
		"https://www.twitter.com",
		"https://www.whatsapp.com",
		"https://www.netflix.com",
	}

	for _, web := range websitelist {
		// getStatusCode(web)
		// go getStatusCode(web) // -> This will not give any output because the main function exits before the goroutine is executed
		// Solution 1: Use time.Sleep() -> This is not a good solution
		// Solution 2: Use WaitGroup -> This is a better solution from package sync
		go getStatusCode(web)
		wg.Add(1) // This is to add a goroutine to the WaitGroup

	}

	wg.Wait() // This is to wait the main function for all the goroutines to finish
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(3 * time.Millisecond) // This is to simulate a delay 
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string){
	defer wg.Done() // This is to remove a goroutine from the WaitGroup -> Responsibility of the developer to mark done
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}else {
		// Locking the resource -> resource is the slice signals
		mut.Lock() // This is to lock the resource
		signals = append(signals, endpoint)
		mut.Unlock() // This is to unlock the resource Here Resource is the slice signals
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}