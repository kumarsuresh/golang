package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait group
	wg := &sync.WaitGroup{}
	//add done or wait functionality

	// mutexes (critical section)
	// if you have a block or db and you want to only one function and go routine can access
	// then you can use mutexes

	// mut := &sync.Mutex()

	wg.Add(2)
	go func ()  {
		fmt.Println("Hello")
		wg.Done()
	}()
	go func () {
		fmt.Println("World")
		wg.Done()
	}()
	
	fmt.Println("Fin")
	wg.Wait()
	fmt.Println("Tin")
	// select{}
}