package main
import (
	"fmt"
)

// All strings in golang are mutable
// means you can change value of string again and again

func main()  {
	var m1 string
	m1 = "Hello World"
	fmt.Println(m1)
}

func main()  {
	//initialize string in one line
	m1 := "Hello World"
	// : is used only for declaration
	// to assign a varible use = sign
	fmt.Println(m1)
}

// if want to check if m2 contains ma or not
func main()  {
	m1 := "Hello World"
	m2 := " !!!"
	// concatenation of string
	fmt.Println(m1 + m2)
	// check m1 contains m2 then it return boolean (True/False)
	fmt.Println(strings.Contains(m1, m2))
	// Replace all o in m1 with yeah
	fmt.Println(strings.ReplaceAll(m1, "o", "Yeah"))
	// split m1 with space and get a array
	fmt.Println(strings.Split(m1, " "))
	// get first index value of array
	fmt.Println(strings.Split(m1, " ")[0])
	// using len calculate lengh of array
	fmt.Println(len(strings.Split(m1, " ")))
}