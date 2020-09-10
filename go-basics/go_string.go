package main

import (
	"fmt"
)

func main() {
	book := "The magic of go"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	// Strings are immutable
	// book[0] = 54

	// Slice (start, end)
	fmt.Println(book[4:11])

	//Slice (no end)
	fmt.Println(book[4:])

	// Slice with no start
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])
	// it support unicode
	// multi line
	poem := `
	Johny johny 
		yes papa
	`
	
	fmt.Println(poem)
}