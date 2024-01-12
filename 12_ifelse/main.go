package main

import "fmt"

func main(){
	fmt.Println("IF-ELSE in Golang");

	loginCount := 23
	var result string;

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "not a regular user, a super user"
	} else {
		result = "normal user"
	}

	fmt.Println(result);

	//another example
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	//in case of web request handling we can check more stuffs in same (if) line
	// like checking a value from the web request that is coming to the system
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is more than 10")
	}

	//to check errors is empty or not
	//if err != nil {
	//	fmt.Println("error found")
	//}

}