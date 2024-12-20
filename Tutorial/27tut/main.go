package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"` // This field will be ignored
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Lets Learn how to handle JSON in Go")
	EncodeJSON()
}

func EncodeJSON() {
	// encoding JSON
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc@123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd@123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit@123", nil},
	}

	// Package this data into JSON
	// finalJSON, error := json.Marshal(lcoCourses)
	finalJSON, error := json.MarshalIndent(lcoCourses, " ", "\t")
	if error != nil {
		panic(error)				
	}
	fmt.Printf("%s\n", finalJSON)}
