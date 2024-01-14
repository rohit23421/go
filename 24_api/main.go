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

//model for api
type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	//creating a type of pointer, referencing the author struct
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

//fake DB of variable name courses and type of course struct
var courses []Course


//middleware, helper
// middleware based on course struct for checking course details empty or not
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}


func main(){
	fmt.Println("API using Golang")
	//router for all the requests
	r := mux.NewRouter()

	//seeding
	courses = append(courses,Course{CourseId: "2", CourseName: "Golang", CoursePrice: 299, Author: &Author{Fullname: "Rohit", Website: "rohit@hi.com"}})
	courses = append(courses,Course{CourseId: "4", CourseName: "C++", CoursePrice: 399, Author: &Author{Fullname: "Rohit sah", Website: "rohit@hiagain.com"}})

	//routing
	r.HandleFunc("/",serverHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",DeleteOneCourse).Methods("DELETE")

	//listening to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}



//controllers

//serve home route that is "/" path
func serverHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome from the API section in Golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r  *http.Request){
	fmt.Println("Get all courses")
	//(w) also helps to set explicit header by headers method and set method
	w.Header().Set("Content-Type", "application/json")
	//in the before line we set a header for json data
	// to through up all the data that is in our DB as json we use json package
	// passing the interface for encode() as the fake DB that we created of format
	// struct for courses
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	//setting the header again for json data type
	w.Header().Set("Content-Type", "application/json")

	//grabbing ID from the request 
	params := mux.Vars(r)
	fmt.Println("Type of params is %T and value is: ",params)
	fmt.Println(params)

	//looping through the range of courses, find matching id and return response
	for _ , course := range courses{
		if course.CourseId == params["id"] {
			//encode and send it through the writer using json package
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// if didnt found any match in the DB
	json.NewEncoder(w).Encode("No course found with the given ID")
	return
}

//for cehcking if the user isnt giving malicous data into the DB
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty means request is empty from user
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//if user is sending data inform of - {}
	// creating a variable of type struct Course
	var course Course
	//we destructure the incoming data
	json.NewDecoder(r.Body).Decode(&course)
	//using the IsEmpty() middleware we created earlier
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No Data inside JSON passed from user")
		return
	}

	//check if title is same or alwys present in DB
	for _ , course := range courses{
		if course.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Same course already exist")
			return
		} else {
			//generating a unique id for course of type string
			// append course into courses slice , that is the fake DB
			rand.Seed(time.Now().UnixNano())
			course.CourseId = strconv.Itoa(rand.Intn(100))
			courses = append(courses,course)
		}
	}

	//generating a unique id for course of type string
	// append course into courses slice , that is the fake DB
	// rand.Seed(time.Now().UnixNano())
	// course.CourseId = strconv.Itoa(rand.Intn(100))
	// courses = append(courses,course)
	json.NewEncoder(w).Encode(course)
	return

}

//updating one course controller
func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update one Course")
	w.Header().Set("Content-Type", "application/json")

	//first looping the DB for finding the course id and then update it
	
	// grab the id from the params using mux
	params := mux.Vars(r)

	//looping throuogh the value and finding the ID and then remove the id
	// from the courses slice(DB) and adding my ID as it is update operation
	// so we need to add the same value but with my ID
	// courses is the DB(slice variable name)
	for index, course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]... )
			//creating a variable of course of type struct Course for passing
			// the reference of the body that i want to decode the request based
			// on this, so we decode using json for the request coming to us
			var course Course // this is the updated value that we want to update
			// inside the fake DB of type Struct Course,
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"] // the new created course have same id
			//even if the user is passing a new id
			courses = append(courses, course) // add the new course to the courses slice
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//sedning a response when id is not found
}

//deleting one course controller
func DeleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	//having a unique id for the coruse to be deleted
	params := mux.Vars(r)

	//looping the DB and if found DB, remove the course using slicing the id
	for index, course := range courses {
		if course.CourseId == params["id"]{
			//updating the fake DB courses slice
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}