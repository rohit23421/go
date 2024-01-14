package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main(){
	fmt.Println("Webserver from Golang")
	PerfromGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

//creating methods for GET request
func PerfromGetRequest(){
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	
	//closing the connection at last of all the lines of code execution
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	//to read all these response use ioutil which gives binary output convert
	// to string as content is in byte format
	//content, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(content))

	//another way of converting the binary to string is using a builder
	var responseString strings.Builder
	contentnew, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(contentnew)

	fmt.Println("Byte count is: ", byteCount)
	// responseString has now a method of reading the byte values directly in
	// form of string as it is derived from the strings.builder method
	fmt.Println("The response is: ", responseString.String())
}

//creating method for POST request for JSON data
func PerformPostJsonRequest(){
	const anotherurl = "http://localhost:8000/post"

	//creating fake json payload(data) using strings package and NewReader method
	requestBody := strings.NewReader(`
		{
			"name": "Rohit Sah",
			"Age": 22,
			"email": "rohit@hi.com"
		}
	`)

	response, err := http.Post(anotherurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content is below \n",string(content))
}

//creating method for POST request for sending FORM_DATA
func PerformPostFormRequest() {
	const anotherurlagain = "http://localhost:8000/postform"

	//creating a fake form patload(data) using the URL package for key:value
	data := url.Values{}
	data.Add("firstname", "rohit")
	data.Add("lastname", "sah")
	data.Add("email", "rohit@hiagain.com")

	response, err := http.PostForm(anotherurlagain, data)
	if err != nil {
		panic(err)
	}

	//when everything done close the request
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content is below \n",string(content))
}