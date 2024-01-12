package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")

	//when the functions go inside structs we call them methods
	rohit := User{"Rohit", "rohit@hi.com", 22, true}
	fmt.Println(rohit)
	fmt.Printf("rohti details are: %+v\n", rohit)
	fmt.Printf("Name is %v and email is %v.\n",rohit.Name, rohit.Email)
	
	//calling the function
	rohit.GetStatus()

	//when we run this , we dont actually change the email value that was 
	//initially passed inside the struct, so a copy of the object was created
	// to actually change this pointers concept was introduced
	rohit.NewMail()
	fmt.Printf("Name is %v and email is %v.\n",rohit.Name, rohit.Email)

}

type User struct {
	Name string
	Email string
	Age int
	IsVerified bool
}

//defining a method, below or above struct doesnt matter
//we have to pass a reeiver inside the method , so we are passing the struct
// we pass a receiver called (u) of type (User) that is a struct
func (u User) GetStatus(){
	fmt.Println("Sending from method")
	fmt.Println("Is user active: ", u.IsVerified)
}

func (u User) NewMail(){
	u.Email = "testing@hi.com"
	fmt.Println("Email of this user is: ", u.Email)
}