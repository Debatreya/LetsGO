package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{}
	// Mutex
	// mut := &sync.Mutex{}
	mut := &sync.RWMutex{}

	// Only Read Lock
	mut.RLock()
	var score = []int{0}
	mut.RUnlock()

	// immediately invoked function expression -> here its a go routine
	wg.Add(3)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("GoR_1")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("GoR_2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("GoR_3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}