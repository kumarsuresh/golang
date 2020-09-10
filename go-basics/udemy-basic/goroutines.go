package main

import (
	"fmt"
	"time"
)

func heavy()  {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}

func super_heavy()  {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super Heavy")
	}
}

func main()  {
	fmt.Println("Hello world")
	go heavy()
	go super_heavy()
	fmt.Println("go routines fin")	
	select {}
}