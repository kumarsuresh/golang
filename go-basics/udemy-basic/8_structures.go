package main

import (
	"fmt"
)


//Structure is a abstract datatype which groups together logically related data
type Car struct {
	Name string
	Age int
	ModelNo int
}

// you can also define the method on top of your structures
func (c Car) Print()  {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("Driving....")
}

func (c Car) GetName() {
	fmt.Println(c.Name)
}

func main() {
	// initialize a variable
	// c := Car{}
	//var c1 Car
	// var is used to declare variable
	// c1 is variable name and
	// Car is datatype and this is called data encapsulation
	// fmt.Println(c)

	// c := Car{"chevy", 1, 2}
	// fmt.Println(c)

	// we can also define a object using object notation
	c := Car{
		Name: "chevy",
		Age: 1,
		// you always suffix last datatype with comma (,)
		ModelNo: 2,
	}
	fmt.Println(c)
	c.Print()
	c.Drive()
	c.GetName()
}

