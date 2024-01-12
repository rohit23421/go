package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	
	//there is no inheretence in Golang and No suepr and parent concepts

	//calling | using the struct created
	Rohit := User{"Rohit", "rohit@hi.com", 22, true}
	fmt.Println(Rohit)
	// %+v is used to extract the value from a struct
	// it gives the values as well as the keys with it from the struct
	fmt.Printf("Rohit's details are: %+v\n", Rohit)
	fmt.Printf("Name is %v and email is %v.\n",Rohit.Name, Rohit.Email)

}

//defining a struct
// starting the inital alphabet as capital letter so that anyone can access
// it becomes a public keyword to be used anywhere
type User struct {
	Name string
	Email string
	Age int
	IsVerified bool
}