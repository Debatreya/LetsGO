package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` // This field will be ignored
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Lets Learn how to handle JSON in Go")
	// EncodeJSON()
	DecodeJSON()
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
	fmt.Printf("%s\n", finalJSON)
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
    {
        "coursename": "ReactJS Bootcamp",
        "Price": 299,
        "website": "LearnCodeOnline.in",
        "tags": ["web-dev", "js"]
    }
	`)

	var lcoCourse course

	// Check of Data is in JSON format
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) // Unmarshal JSON data into Go data, &lcoCourse is the address of the variable
		fmt.Printf("%#v\n", lcoCourse) // %#v is used to print the struct
	}else {
		fmt.Println("JSON was not Valid")
	}

	// some cases where you just want to add data to key value pair
	var myOnlineData map[string]interface{} // map with key as string and value as interface (can be anything) , we dont know the type of value
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is: %v and Value is: %v and Type is: %T\n", k, v, v)
	}
}
