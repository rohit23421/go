package main

import "fmt"

func main(){
	fmt.Println("Maps in Golang");

	//map is like key: value pairs
	//creating a map of key = string and value as string
	languages := make(map[string]string)

	languages["C"] = "C language"
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	//we can retrive a single vaalue from the maps like using the key
	fmt.Println("JS shorts for: ", languages["JS"])

	//delete a value form the Map
	delete(languages, "PY")
	fmt.Println("New list of languages: ",languages)

	//loops in Golang
	for key, value := range languages {
		// %v is for value 
		fmt.Printf("For key: %v, value is %v\n", key, value)
	}


}