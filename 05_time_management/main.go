package main

import (
	"fmt"
	"time"
)

func main() {
	
	//using time library for handling time in golang
	fmt.Println("welcome to 5th chapter for time management in GO");

	presentTime := time.Now();
	fmt.Println(presentTime)
	// always use these specific numbes to fetch the current dte and time & day
	// using format method to filter spcific values
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	//creating a date and time using Date method
	createdDate := time.Date(2023, time.December, 19, 12, 0, 1, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}