package main //main package make your code compile and execute

import (
	"fmt" //fmt is for formating printed line and logging
)

// array is list of ordered items

func todo() {
	// Declaring and initializing array
	// var arr = []int
	arr := []int{1,2,3,4}
	fmt.Println(arr)

	// you also crete a string array
	arr2 := []string{"hello", "world"}
	fmt.Println(arr2)

	// we use append function to append things in array
	// append one item
	arr2 = append(arr2, "xyz")
	fmt.Println(arr2)

	// append multiple items
	arr2 = append(arr2, "abc", "hi", "hello")
	fmt.Println(arr2)

	// get array length
	fmt.Println(len(arr2))
	fmt.Println(arr2[4])
	fmt.Println(arr2[:3])
	fmt.Println(arr2[3:4])

	// Go doesn't have negative indexing like Python does
	fmt.Println(arr2[-1])

	fmt.Println(arr, arr2, "Sa re ga ma")
}
// Every file has a main function which is a entry point of every file
func main() {
	todo()
}