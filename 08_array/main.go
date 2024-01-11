package main

import "fmt"

func main() {
	fmt.Println("Hello there this is about arrays in Golang");

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[2] = "Banana"
	fruitList[3] = "Pinnaple"

	//if we miss a index in the arrat and print it we can see that there will
	// be a space printed in output to show that there is one more space of memory
	fmt.Println("Fruit list is: ",fruitList)

	//if we do length in this rray it will give the size of total array either if
	// the array has only 3 items istead of 4
	fmt.Println("length of fruitList is: ",len(fruitList))

	var vegList = [3]string{"potato","beans","mushroom"}
	fmt.Println("veg list is: ",vegList[1]);
	fmt.Println("length of veg list is: ",len(vegList));
}