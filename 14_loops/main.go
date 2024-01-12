package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	//creating a slice here for looping
	days := []string{ "Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for day:=0; day < len(days); day++ {
		fmt.Println("Day is: ", days[day])
	}

	//another way of finding the size of the slice
	for d := range days {
		fmt.Println(days[d])
	}

	//another way using comma ok method
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v \n", index, day)
	}

	//break and continue
	somevalue := 1
	for somevalue < 10 {

		//using the goto statement to catch the label name (rohit)
		if somevalue == 2{
			goto rohit
		}

		//use either continue or break
		if somevalue == 5 {
			somevalue++
			continue
		}
		if somevalue == 5 {
			break
		}

		fmt.Println("Value is: ", somevalue)
		somevalue++
	}

	//goto statement in Golang here by using a labelname for calling in above code
	rohit:
		fmt.Println("Testing the Goto statement in Golang")

}