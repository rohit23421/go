package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//gloabal varibale where we define the URL
const url = "https://lco.dev"

func main(){
	fmt.Println("Web Requests in Golang")

	//making the http request using the http package
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of response is %T \n", response)

	//caller's responsidbility is to close the connection at last after calling
	defer response.Body.Close()

	//to read the response that we received
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println(content)
}