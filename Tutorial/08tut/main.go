package main

import (
	"fmt"
	"time"
)

func main() {
	// This is a standalone program that will give me the DATE-TIME-DAY whenever it is run
	// I will use the time package to get the current time

	currentTime := time.Now()

	// I will use the Format method to format the time in a human readable format

	fmt.Println("Today is", currentTime.Format("2006-01-02 Monday"))
}
