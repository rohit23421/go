package main

import "fmt"

//can create a const value
// the first letter is captial means its a public value and can be used anywhere
const LoginToken string = "hdghdgkdh"


// in GO when we declare a variables we have to use it or else it will throw error
func main() {
	var username string = "Rohit";
	fmt.Println(username);
	// %T is used to extract the type of variable
	fmt.Printf("Variable is of type: %T \n", username);

	var isLoggedIn bool = false;
	fmt.Println(isLoggedIn);
	fmt.Printf("variable is of type: %T \n", isLoggedIn);

	var smallVal uint8 = 255
	fmt.Println(smallVal);
	fmt.Printf("variable is of type %T \n", smallVal);

	var smallfloat float32 = 2434.33536786563557843;
	fmt.Println(smallfloat);
	fmt.Printf("variable is of type: %T \n", smallfloat);

	//default values and some aliases
	var anotehrvar int;
	fmt.Println(anotehrvar);
	fmt.Printf("variable is of type %T \n", anotehrvar);

	//implicit type of declaring a varibale
	// in GO as we know we have to statially define a variable but here lexer
	// automatically does it for us even if we dont say the website variable is
	// a string variable but we cant further make changes to the same variable
	// we can assign a differnet type of variable like int lexer doesnt allwos it
	var website = "google.com";
	fmt.Println(website);
	//website = 3;

	//no var style variable where we dont write the var and vsriable keyword
	numberofUser := 300000.0 // this (:=) walrus method can only be used inside a functions method, cannot be used outside func main()
	fmt.Println(numberofUser)

	//using the public variable declared at the top
	fmt.Println(LoginToken);
	fmt.Printf("variable type is of: %T \n", LoginToken);
}