package main

import (
	"fmt"
)

func main() {
	flag := true
	// continue, break, switch case
	for i := 0; i < 10; i ++ {
		if i%2 == 0 {
			// continue
			if i == 6 {
				flag = false
				break
			} else {
				continue
			}
		}
		// fmt.Println(i)
	}
	fmt.Println(flag)


	//Switch example
	day := "Fri"
	switch day {
		case "Mon":
			fmt.Println("Hangover")
		case "Fri", "Sat", "Sun":
			fmt.Println("Party")
		default:
			fmt.Println("Working")
	}

	switch  {
		case day == "Fri":
			fmt.Print("dk")
			break
		
	}
}