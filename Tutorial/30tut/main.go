package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct { // Capitalized - public, can be exported
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice float64 `json:"price"` // "-" - ignore this field while encoding to JSON, then what key to use in case of POST request? -> use the field name
	Author      *Author `json:"author"` // Pointer to Author struct, cant use '&' operator here as it is not a pointer
}


type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool { // Method on Course struct defined using (c *Course) receiver
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to the course API")
	r := mux.NewRouter() // create a new router


	// seeding the fake DB
	courses = append(courses, 
		Course{
			CourseId:    "1",
			CourseName:  "Go Programming",
			CoursePrice: 100.00,
			Author: &Author{
				FullName: "John Doe",
				Website:  "https://johndoe.com",
			},
		},
		Course{
			CourseId:    "2",
			CourseName:  "Python Programming",
			CoursePrice: 120.00,
			Author: &Author{
				FullName: "Jane Doe",
				Website:  "https://janedoe.com",
			},
		},
	)

	// routes
	r.HandleFunc("/", serveHome).Methods("GET") // home route
	r.HandleFunc("/courses", getAllCourses).Methods("GET") // get all courses route
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET") // get course by id route
	r.HandleFunc("/courses", createOneCourse).Methods("POST") // create one course route
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT") // update one course route
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE") // delete one course route

	// listening to a port
	log.Fatal((http.ListenAndServe(":4000", r))) // listen to the port 4000
}

// controller - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the home page API</h1>"))
}

// get all courses route
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set(("Content-Type"), "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get course by id route
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set(("Content-Type"), "application/json")

	// grab the id from the request
	params := mux.Vars(r) // get the params from the request using mux
	// loop through the courses and find the course with the id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course with given ID")
}

// create one course route
func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create one Course")
	w.Header().Set(("Content-Type"), "application/json")

	// What if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	// what about - {}
	var course Course
	// decode the request body into course
	_ = json.NewDecoder(r.Body).Decode(&course) // decode the request body into course
	// check if the course is empty
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}
	// generate a random id, string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(1000)) // generate a random number between 0 and 1000
	// append the course to the courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

// update one course route
func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one Course")
	w.Header().Set(("Content-Type"), "application/json")

	// first - grab the id from the request
	params := mux.Vars(r) // get the params from the request using mux // params['id'] - id is the key
	// loop, id, remove, add with my ID

	for index, course := range courses {
		// Match the course with the id with params['id']
		if course.CourseId == params["id"] {
			prevCourse := courses[index] // store the previous course
			courses = append(courses[:index], courses[index+1:]...) // remove the course from the courses
			var updateInCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updateInCourse) // decode the request body into course
			updateInCourse.CourseId = params["id"]
			// if the course name is empty, use the previous course name
			if updateInCourse.CourseName == "" {
				updateInCourse.CourseName = prevCourse.CourseName
			}
			// if the course price is 0 or not present, use the previous course price
			if updateInCourse.CoursePrice == 0 {
				updateInCourse.CoursePrice = prevCourse.CoursePrice
			}
			// if the author is nil, use the previous author
			if updateInCourse.Author == nil {
				updateInCourse.Author = prevCourse.Author
			}
			// append the course to the courses
			courses = append(courses, updateInCourse)
			json.NewEncoder(w).Encode(updateInCourse)
			return
		} // loop through the courses
		
	} // loop through the courses
	
	// Send a response if no course with the given ID
	json.NewEncoder(w).Encode("No Course with given ID")
}

// delete one course route
func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one Course")
	w.Header().Set(("Content-Type"), "application/json")

	// grab the id from the request
	params := mux.Vars(r) // get the params from the request using mux
	// loop through the courses and find the course with the id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) // remove the course from the courses
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No Course with given ID")
}	