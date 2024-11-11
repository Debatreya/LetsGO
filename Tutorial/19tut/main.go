package main

import "fmt"

// Declaring Struct (outside main function)
type User struct { // User is a struct, U - capital is for public access same for fields, if we want to make it private, we can use small u
	Name       string
	Email      string
	isVerified bool
	Age        int
	status     string
}

func main() {
	fmt.Println("Methods in GoLang")
	user2 := User{
		Name:       "Jane Doe",
		Email:      "janedoegmail.com",
		isVerified: true,
		Age:        30,
		status:     "Active",
	}
	fmt.Println(user2)
	// Calling a Method
	user2.NewMail()
	// The above method call will not change the email of user2, because the method is not changing the original value of the struct
	// To change the original value of the struct, we need to use pointers
	// This is because Go is a pass by value language for methods
	fmt.Printf("For USER2: Name is %v and Email is %v\n", user2.Name, user2.Email) // Accessing struct fields

	// Calling a Method
	user2.GetStatus()
}


// Declaring a Method
func (u User) GetStatus(){
	fmt.Println("Is User Active: ", u.status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("New Email is: ", u.Email)
}