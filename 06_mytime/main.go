package main

import (
	"fmt"
	"time"
)

//after making this GO file do a go build in the cmd for windows system
// GOOS=windows go build | for linux | GOOS=linux go build
// an .exe file will be generated in the directory for running this file

func main() {
	presentTime := time.Now();
	//fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}