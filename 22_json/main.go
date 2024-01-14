package main

import (
	"encoding/json"
	"fmt"
)

// we can use the json help for call the variables of our struct differently
// when sending the request like calling Name as username
type user struct {
	Name string `json:"username"`
	Age int
	Email string `json:"emailId"`
	Password string `json:"-"` // (-) means, we dont want to reflect the password
	Tags []string `json:"tags,omitempty"` // omitempty work so as if there is
				// any nil/null data, it omits it instead of sending them as null
}

func main(){
	fmt.Println("JSON response in Golang")
	EncodeJson()
	DecodeJson()
}

//method for creating json data
func EncodeJson(){

	userpresent := []user{
	{"Rohit sah", 22, "rohti@hi.com", "abc123", []string{"normal user", "prime"}},
	{"Foo barh", 22, "rohti@hi.com", "bcd123", []string{"new user", "prime"}},
	{"John doe", 22, "rohti@hi.com", "dsgs123", nil},
	}

	//package this data as jSON data
	// using the json package to create the json data and indext them after ""
	finalJson, err := json.MarshalIndent(userpresent, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

//method for consuming JSON data
func DecodeJson(){
	//as the data that comes from web is always in byte format, so we need to
	// convert it to be consumed inform of json data
	jsonDataFromWeb := []byte(`
	{
		"username": "Rohit sah",
        "Age": 22,
        "emailId": "rohti@hi.com",
	    "tags": ["normal user","prime"]
	}
	`)

	// creating object from struct
	var username user

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		//unmarshal takes the byte data and interface, in our case user (struct)
		json.Unmarshal(jsonDataFromWeb, &username)
		fmt.Printf("%#v\n", username)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some cases where we just want to add data to key value
	// so we sometimes can use a map
	// when we reate a map for web request the key can be string but the incoming
	// data that is coming can be of any type so we kept it as interface
	var myOnlineData map[string]interface{}
	// decoding the web data first using unmarshal from json package by passing
	// the jsondatafromweb and the actual address of the map variable created
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("Online data is below \n %#v\n",myOnlineData)

	//another way to fetch this map is
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n",k,v,v)
	}
}