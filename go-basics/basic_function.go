// Go comments gone here
package main //main package make your code compile and execute

import (
	"fmt" //fmt is for formating printed function
)

func add(x int, y int) int  {
	return x + y
}

func divmult(x float64, y float64) (float64, float64)  {
	return x/y, x*y
}

func main() {
	val := add(13, 2)
	fmt.Println(val)

	val1, val2 := divmult(13.0, 2.0)
	fmt.Printf("%v - %v %T", val1, val2, val1)
}

// go test => have build in test cases
// go get => to install third party packages
