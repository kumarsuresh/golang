package main

import (
	"fmt"
)

func main() {
	for i :=0; i<5; i++ {
		if i > 3 && i < 5 {
			break
		} 
		
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("-----------------")
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Printf("-----------------")
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}


}