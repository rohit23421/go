package main

import "fmt"

func main() {
	fmt.Println("Pointers in GO")

	// if we declare a pointer without any assignment and print it we get the 
	// default value as <nil>
	var pointer *int
	fmt.Println("value of pointer is: ", pointer)

	myNumber := 23;
	// using (&) to reference the pointer that reference to the menmory location
	var ptr = &myNumber
	fmt.Println("value of actual pointer memory location is: ", ptr)
	fmt.Println("value of actual pointer is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("new value of the actual pointer is: ", myNumber);
}