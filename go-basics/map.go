// Go comments gone here
package main //main package make your code compile and execute

import (
	"fmt" //fmt is for formating printed function
)

func main() {
	fmt.Printf("hello, world\n")
	_stocks := map[string]float64{
		"AMN": 234.44,
		"GOOG": 783.33,
		"TSL":263.14,
	}
	fmt.Println(_stocks)
	fmt.Println(len(_stocks))
	fmt.Println(_stocks["TSL"])
	fmt.Println(_stocks["HELLO"])
	value, ok := _stocks["HELLO"]

	if !ok {
		fmt.Println("HELLO is not found")
	} else {
		fmt.Println(value)
	}


	fmt.Println("-----------------------")
	_stocks["PILLOW"] = 978.23
	fmt.Println(_stocks)
	fmt.Println(_stocks["PILLOW"])
	delete(_stocks, "PILLOW")
	fmt.Println(_stocks["PILLOW"])
	fmt.Println("-----------------------")

	for key := range _stocks {
		fmt.Println(key)
	}
	fmt.Println("-----------------------")

	for key, value := range _stocks {
		fmt.Println(key)
		fmt.Println(value)
	}
}

// go test => have build in test cases
// go get => to install third party packages
