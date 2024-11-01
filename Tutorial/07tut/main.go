package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	// time package
	// time package is used to work with time in golang
	// time package provides time.Time type which represents time

	// time.Now()
	presentTime := time.Now()
	fmt.Println("Present time is:", presentTime)

	// formatting time
	// time.Format(layout string) string
	// layout is a string which represents how the time should be formatted
	// layout is a reference to the time package documentation
	// https://golang.org/pkg/time/#pkg-constants
	fmt.Println("Present Date is:", presentTime.Format("01-02-2006")) // 01-02-2006 is a reference to the time package documentation and it is a constant which represents the layout of the date
	fmt.Println("Present time is:", presentTime.Format("01-02-2006 15:04:05")) // 01-02-2006 15:04:05 is a reference to the time package documentation and it is a constant which represents the layout of the date and time
	fmt.Println("Present time and Day is:", presentTime.Format("01-02-2006 15:04:05 Monday")) // 01-02-2006 15:04:05 Monday is a reference to the time package documentation and it is a constant which represents the layout of the date and time and Day

	// Create a Date
	createdDate := time.Date(2029, time.November, 27, 0, 0, 0, 0, time.UTC)
	fmt.Println("Created Date is:", createdDate.Format("01-02-2006 15:04:05"))
	// Print the Day of the created date
	fmt.Println("Day of the created date is:", createdDate.Weekday())
}
