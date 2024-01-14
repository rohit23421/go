package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=3423hjh43"

func main() {
	fmt.Println("Url in Golang")
	fmt.Println(myurl)

	//parsing the url to extract information that we need
	// by url library we can parse the url
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	//scheme here is https so we can get it from the results
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port()) // port is a method, not a property
	fmt.Println(result.RawQuery) // all the queries that was passed to the url

	// saving the details we want into a variable
	qparams := result.Query() // this Query() method saves the data in a better format
	fmt.Printf("The type of query params are: %T\n", qparams)
	// the type of query() is url.values that is same as key:alue pair
	// so we can use the key as coursename to fetch the value from url
	fmt.Println(qparams["coursename"])

	//extracting the values of key:value pairs
	for _,val := range qparams {
		fmt.Println("Param is: ",val)
	}

	//when we do this method we only pass the refernece of the url package
	//not the copy so we use the & sign to signify the refernce
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=rohit",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)


}