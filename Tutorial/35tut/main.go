package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	// Declaring a channel
	// myCh := make(chan int) // myCh is a channel of type int
	// Buffer channel
	myCh := make(chan int, 2) // myCh is a channel of type int , 2 is the buffer size


	// Sending and receiving values from the channel // But this is a Deadlock
	//// myCh <- 23 // sending 23 to the channel
	//// fmt.Println(<-myCh) // receiving the value from the channel

	// Deadlock because we are sending a value to the channel but there is no receiver to receive the value
	// So, the program will be stuck at line 11 and will not proceed further

	// To avoid deadlock, we can use goroutines
	wg := sync.WaitGroup{}
	wg.Add(2)

	// This will not cause deadlock
	// This goroutine will receive the value from the channel
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		if(isChannelOpen){
			fmt.Println(val)
		} else {
			fmt.Println("Channel is closed", val)
		}
		wg.Done()
	}(myCh, &wg)

	// This goroutine will send the value to the channel
	go func(ch chan int, wg *sync.WaitGroup) {
		// myCh <- 0
		myCh <- 24
		// myCh <- 25
		close(myCh) // Closing the channel
		wg.Done()
	}(myCh, &wg)

	wg.Wait()
}