package main

import "fmt"

// Declaring Struct (outside main function)
type User struct { // User is a struct, U - capital is for public access
	Name       string
	Email      string
	isVerified bool
	Age        int
}

func main() {
	fmt.Println("Structs in GoLang")
	// no inheritance in GoLang; No classes in GoLang; No Super ... in GoLang
	// GoLang has structs

	// Structs are value types
	// Structs are value types and are copied when passed around

	var user1 User = User{"John Doe", "johndoegmail.com", false, 25} // Order of fields should be same as in struct
	// We can also declare struct like this
	// var user1 User = User{Name: "John Doe", Email: "johndoegmail.com", isVerified: false, Age: 25} // Order of fields can be different
	fmt.Println(user1)

	// Other way to declare struct (using walrus operator)
	user2 := User{
		Name:       "Jane Doe",
		Email:      "janedoegmail.com",
		isVerified: true,
		Age:        30,
	}
	fmt.Println(user2)
	fmt.Printf("user2 is %+v\n", user2) // %+v will print field names as well
	
	// Accessing struct fields
	fmt.Println("Name of User user1 is", user1.Name)
	fmt.Printf("For USER2: Name is %v and Email is %v\n", user2.Name, user2.Email) // Accessing struct fields

	// We will get into it with more detail later with the APIs
}
