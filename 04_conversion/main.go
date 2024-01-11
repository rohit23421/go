package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our store")
	fmt.Println("Please rate our products out of 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating the product a: ", input)

	//do something to the input value rating
	// convert the string input from user to integer using str.conv package
	// and we want to conver it to a float so parsefloat is used
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err);
		//we can also add a panic and pass the error to stop the program
		// panic(err)
	} else {
		fmt.Println("Added 1 to the rating: ",numRating+1)
	}
}