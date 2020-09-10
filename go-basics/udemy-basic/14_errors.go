package main

import (
	"fmt"
	"math"
)

//go function can return more then one value
// this we use to signal errors

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0.0, fmt.Errorf("Sqrt of negative value %f", n)
	}
	return math.Sqrt(n), nil
}

func main() {
	s1, err := sqrt(2.0)

	if err != nil {
		fmt.Printf("Error %s\n", err)
	} else {
		fmt.Println(s1)
	}

	s2, err2 := sqrt(-2.0)

	if err2 != nil {
		fmt.Printf("Error %s\n", err2)
	} else {
		fmt.Println(s2)
	}
}