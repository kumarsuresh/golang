/**
channel is just a queue with mutexes
Channel is or multiple type
buffered channel  (queue size is dynamically)
unbuffered channel (queue size is 1)
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// var c chan int
	c := make(chan int)

	// send in a goroutine
	go func() {
		c <- 12
	}()
	val := <-c
	fmt.Println(val)
	go func() {
		c <- 1
	}()
	// sniff
	time.Sleep(time.Second * 2)
	val = <-c
	fmt.Println(val)
	fmt.Println(c)
}