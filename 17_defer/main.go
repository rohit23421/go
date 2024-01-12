package main

import "fmt"

func main(){
	fmt.Println("Defer in Golang")

	//A defer statement defers the execution of a function until the surrounding
	//function returns The deferred call's arguments are evaluated immediately
	//but the function call is not executed until the surrounding function returns

	defer fmt.Println("World") // this wont print before Hello as its defer
	fmt.Println("Hello")
	//deferred statement exeutes at the last of the funcion execution

	//defer works in LIFO principle
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three") // output will be => three | two | one

	//calling the deferred function to see the flow as the function call will
	// happen first before rpinting the deeferred lines in this funciton
	//so myDefer() vlues will be printed first then statements from this function
	// Hello | 4 | 3 | 2 | 1 | 0 | three | two | one | world
	myDefer()
}

func myDefer(){
	for i:=0; i <5; i++ {
		defer fmt.Println(i)
	}
}