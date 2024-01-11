package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input";
	fmt.Println(welcome);

	//storing the variable value using the Buff.io package
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter a rating for the product: ");

	//storing the read value from the reader to a variable
	// in GO we dont have a try catch for treating vslues and errors
	// we used ( comma ok || err ok ) method

	//here we can do extract the input in a variable whatever we want to call it
	// be it a singel alphabet or word, same goes for error we can store the error
	// variable or in GO we simple do a ( _ )
	input, _  := reader.ReadString('\n') //read upto the when the line is changed
	fmt.Println("Thanks for rating the product a: ", input);
	fmt.Printf("Type of rating the product is: %T", input);
	
}