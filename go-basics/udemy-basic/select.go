package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quits chan struct{}) {
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("received")
		case <- quits:
			fmt.Println("Quit")
			os.Exit(1)
		}
	}
}

func main()  {
	fmt.Println("Hello world")
	c := make(chan int, 2)
	quits := make(chan struct{})
	
	go func ()  {
		c <- 1
		quits <- struct{}{}
	}()

	go Select(c, quits)
	select{}
}