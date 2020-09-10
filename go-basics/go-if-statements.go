package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 && x < 10 {
		fmt.Printf("x is big")
	} else {
		fmt.Printf("x is not big")
	}

	a := 11.0
	b := 20.0

	if frac := a/b; frac > 0.5 {
		fmt.Printf("a is more than half of b")
	}
}