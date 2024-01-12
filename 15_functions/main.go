package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greeter();

	//we cant delcare a function inside a function in Golang,we have to do outside
	// func greetertwo() {
	// 	fmt.Println("Another greeting from Golang")
	// }
	// greetertwo()

	result := adder(3,5)
	fmt.Println("Result is: ", result)

	advanceresult, advancemessage := advanceadder(2,4,6,9,1)
	fmt.Println("Advance result is: ", advanceresult)
	fmt.Println("Advance result is: ", advancemessage)
}

// we have to define function signature for what kind of value to pass and 
//what kind of value to return, here we pass ints and return (int)
func adder(val1 int, val2 int) int {
	return val1 + val2
}

// this is how we added a number of vlues, if we dont know how many values are
// going to be passed, so we added all of them when we receive it so we use (...)
func advanceadder(values ...int) (int, string ) {
	total := 0

	for _,val := range values {
		total += val
	}
	return total, "This is coming from advanceadder"
}

func greeter(){
	fmt.Println("Hello from Golang")
}