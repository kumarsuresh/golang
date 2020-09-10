package main

import (
	"fmt"
)

func main() {
	x := 2
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Many")
	}

	switch {
	case x > 100:
		fmt.Println("bigger then 100")
	case x > 1:
		fmt.Println("Bigger then 1")
	default:
		fmt.Println("x is small")
	}
}