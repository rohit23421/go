package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Hello there from the slices section in Golang");

	//when we want to use slices we need to initialize the array at the declaration
	var fruitList = []string{"Apple","Banana","Peach"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)
	
	//to append/add these values to fruitlist which was dynamic size at declartion
	fruitList = append(fruitList, "Mango", "Tomato")
	fmt.Println(fruitList)

	//to slice from 1st index position of array to end of array leave the 0th index
	fruitList = append(fruitList[1:])
	//ranges are non inclusize of the last value that is upto value
	var fruitListtwo =  append(fruitList[1:3])
	fmt.Println(fruitList)
	fmt.Println(fruitListtwo)

	//making a array with size 4 of integer type
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	fmt.Println(highScores)
	//although we declare the array with 4 size, if we again append new values
	// GO does teh memory allocatio again with the new size
	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)

	//sorting in GO to help sort string/interger/float/ boolean answers for sorted or not
	// sort the values of integers in ascending order
	sort.Ints(highScores)
	fmt.Println(highScores)
	//if move this line before sorting at line 42 then we get the output as false
	fmt.Println(sort.IntsAreSorted((highScores)))

	//removing a value from slices based on index
	var courses = []string{"reactjs","java","aws","python","ruby"}
	fmt.Println("courses are: ", courses)
	var index int = 2; // index value to be deleted and last value is inclusive
	//we can also use append to remove a speific index from the slice
	// [:index] means I dont have a first value so start from beginneing and go till
	// index value and then second value is [index+1] means 2+1=3 and go till end
	// means we append in two halfs here first from index 0 to 1st then 3rd to last
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}