package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Mods in Golang")
	greeter()

}

func greeter(){
	fmt.Println("Hey there mod users")
	//mux is a router for Golang
	r := mux.NewRouter()
	// routing to "/" path for GET request
	r.HandleFunc("/", serverHome).Methods("GET")

	//running it as a server
	// using Log package for handling the error
	log.Fatal(http.ListenAndServe(":4000", r))
}

// making a web request and hadnling the response
// where the request is a pointer 
func serverHome(w http.ResponseWriter, r *http.Request) {
// if we want to send some response we have to use (w) as (r) is someone sending
// us a request so we have to response using (w) 
	//we are sending using the (W) to write a response and on web its in byte
	// format so we convert the string to byte
	w.Write([]byte("<h1>Welcome to Golang modules section</h1>"))
}