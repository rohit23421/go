package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Files in Golang")

	//we want to put this line into a file
	content := "This line needs to go into a file"

	//we use os library for creating a file into the current working directory
	file, err := os.Create("./testfile.txt")
	if err != nil {
		//panic will shutdonw execution of program and show us the error
		panic(err)
	}
	//if everythings works fine we want to write the line to a file
	// by using the io package
	// it takes the file and the string value to be written inside it
	// if it executes properly, it returns us the length of the file
	length, err	:= io.WriteString(file, content)
	//use the function created at last for avoid writing same line of code
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	// now after successful attempt we want the file to close
	// we used defer so that it always executes at the last of theprogram
	defer file.Close()
	readFile("./testfile.txt")
}

//creating another function to read the file that was written
func readFile(filename string){
	//for reading the file we have to use (ioutil) package and for writing (os)
	//whenever we read a file we dont read in string, we read it in byte format
	databyte, err := ioutil.ReadFile(filename) // byte format called databyte
	checkNilErr(err)

	fmt.Println("Text data inside the file is: \n",databyte)
	//converting the databyte data to string
	fmt.Println("Text data inside the file is: \n",string(databyte))

}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}