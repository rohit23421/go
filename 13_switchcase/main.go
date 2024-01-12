package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case - Dice game")

	//using math.random package to generate a random dice value
	// here seed is the input from where to generate a random number
	// so for randomness we passed time, we can pass anything and then for more
	// precise randomness we used time.Now() and UnixNano()
	rand.Seed(time.Now().UnixNano())
	// here we are checking the dicenumber should be in (Intn) 6, exlusive of 6
	// so we added 6+1 that now becomes 6 and then store it in a variable
	dicenumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open now")
	case 2:
		fmt.Println("you can move to 2 spot")
	case 3:
		fmt.Println("you can move to 3 spot")
		//fallthrough works same like continue in C/c++ languages
		fallthrough
	case 4:
		fmt.Println("you can move to 4 spot")
		fallthrough
	case 5:
		fmt.Println("you can move to 5 spot")
	case 6:
		fmt.Println("you can move to 6 spot")
	}

}