// Go comments gone here
package main //main package make your code compile and execute

import (
	"fmt" //fmt is for formating printed function
	"strings"
)

func main() {
	text := `
		Needles and pins 
		Needles and pins
		sew me a sail
		to catch me the wind
	`
	fmt.Println(text)
	fmt.Println("------------------")
	words := strings.Fields(text)
	fmt.Println(words)
	counts := map[string]int{}
	fmt.Println(counts)
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)
}

// go test => have build in test cases
// go get => to install third party packages
