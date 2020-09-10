package main

import (
	"fmt"
)
// Interfaces are named collection of method signatures
type Car interface {
	// you can define certain sets of function which you want other classes should follow
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

// we can define some Drive and Stop function on top of these structs
// so that they start following the interface
func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main() {
	l := Lambo{"Galardo"}
	c := Chevy{"C387"}
	l.Drive()
	c.Drive()
}

