package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channles in Golang")
	// channels ar a way so that goroutine can commutnicate with other
	// not exactly communicate, it means it makes go rotuines can wait for each
	// others process to finish, altough they may not be able to know what
	//exactly the other go routine is doing
	//channles is pipleine that we have to define taht what kind of data
	// the gorotuines to be passing to each other
	// we are pass (2) as two listener when creating the channel so that
	//two values can be passed
	myCh := make(chan int, 2)
	//waitgroup pointer
	wg := &sync.WaitGroup{}

	//fmt.Println(<-myCh)
	//myCh <- 5

	//goroutines
	wg.Add(2)
	// read only rotuine // value going outside channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		//recieve the value here / recive the value outof channel
		fmt.Println(<-myCh)
		fmt.Println(<-myCh)
		//we can also recieve the value in a variable, only if channel is open
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	//send only routine // value coming inside channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		//myCh <- 5 // write to the channel / into the channel
		//to pas two value we have to have two listenner in teh other routine
		// while using channel because, in channel lister is neccessary
		//myCh <- 6
		close(myCh) // closing the channel after passing the values
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
