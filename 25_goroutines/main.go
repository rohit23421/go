package main

import (
	"fmt"
	"net/http"
	"sync"
	//"time"
)

// we make the thread to work and wait for its repsonse for 2 sec then it sends
// us the result this is gorutine synchronization

//a global variable for a wait group ,but in majority cases we make a pointer
// using sync pakcage for go routine
var wg sync.WaitGroup // usually these are pinters

//mutex provides a lock over a memomry till the time one goroutine is working
// to write on that memor spcae
var mut sync.Mutex // usaully a pointer

//simple string array varable for appending the written data adfter lokcing with
//mutex to be read and appended in last line of code
var signals = []string{"test"}

func main() {
	fmt.Println("Goroutines in GOLANG")
	//create a go routine for only hello
	// this means fire up a thread to print hello and wait till it comes back
	// and respnods us
	//go greeter("Hello")
	//greeter("World")

	//wait group
	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	//get status code from websitelist and passon the value to the istance
	for _, web := range websiteList {
		go getStatusCode(web)
		//adding the waitgroup add that will do parallelism for process
		wg.Add(1)
	}

	//this makes sure that the main method doesnt get close before getting all
	//done, means asking main method to wait till all response comes back
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		// passing the time for goroutine to sleep for some time adn come back
// 		time.Sleep(2 * time.Second) // this makes the loop to wait 2 second tehn print
// 		fmt.Println(s)
// 	}
// }

//wait groups in golang
func getStatusCode(endpoint string) {

	//wait group defer for making sure all the response has came back before exiting
	// this funtion and asking main func to wait till all response
	defer wg.Done()

	fmt.Println("Inside Statuscode")
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error in getStatuscode endpoint")
	} else {
		//before writing on menory block ask mutex to lock that space first
		mut.Lock()
		signals = append(signals, endpoint)
		//after writing on that space unlock the space for read purpose
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}
}
