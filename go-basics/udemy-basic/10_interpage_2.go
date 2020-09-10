package main

import (
	"fmt"
)

// an empty interface is a black box it can be int/string/float anything
func Anything(anything interface{})  {
	fmt.Println(anything)
}

type MyName struct {
	Name string
}
func main() {
	fmt.Println("here we go")
	Anything(2.44)
	Anything("Suresh")
	Anything(struct{}{})
	Anything(MyName{"Kumar"})

	mymap := make(map[string]interface{})
	mymap["name"] = "Suresh"
	mymap["Age"] = 23
	mymap["height"]=57.5
	fmt.Println(mymap)
}