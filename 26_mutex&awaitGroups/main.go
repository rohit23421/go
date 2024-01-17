package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition in Golang")

	//creating a pointer wait group
	wg := &sync.WaitGroup{}
	//creating a mutex for locking and unlocking and using RACE for Read/write
	mut := &sync.RWMutex{}

	//a socre variable to store values with default as 0, more value will be pushed
	var score = []int{0}

	//some goroutinee
	// we can do  this add 3 times for each routine to make the waitgroup know
	// to add one more routone or there is an alternative for thsi also
	//wg.Add(1)
	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("First routine")
		//using the mutex to lock while we perform write
		mut.Lock()
		score = append(score, 1)
		//unlokcing after writing on the memory space
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	//wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Second Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	//wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Third Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//if we just use Mutex(not RWmutex) we dont need to add Rlock(Read Lock)
	// Read is a quick operaiton compared to write
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Fourth Routine")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
