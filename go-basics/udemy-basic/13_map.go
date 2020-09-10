package main

import (
	"fmt"
)

func main()  {
	stocks := map[string]float64 {
		"amzn": 234.2,
		"gogl": 8723.2,
		"mcft": 342.5,
	}

	// nUmber of items
	fmt.Println(len(stocks))

	//get specific value
	fmt.Println(stocks["amzn"])

	//Get zero value if key not found
	fmt.Println(stocks["tsla"])

	// check if key is exist or not
	value, ok := stocks["tsla"]
	if !ok {
		fmt.Println("key not found")
	} else {
		fmt.Println("key found")
		fmt.Println(value)
	}

	// set value
	stocks["tsla"] = 239.2
	fmt.Println(stocks)

	// Delete
	delete(stocks, "amzn")
	fmt.Println(stocks)

	// iterate over map
	for key, value := range stocks {
		fmt.Println(key)
		fmt.Println(value)
	}
}