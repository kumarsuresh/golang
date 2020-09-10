// Gonna be talking about pointers
// a pointer is just a variable which stores the address of another variables.package basic
package main

import (
	"fmt"
)

func main() {
	m1 := 20
	ptr := &m1
	//& symbol is refrencing symbol

	// it prints and address of ptr
	fmt.Printf("Pointer %v\n", ptr)
	fmt.Println(ptr)

	// we have to de-refrence the variable to get exact value of ptr
	// we use * to de-refrence
	fmt.Printf("Pointer %v\n", *ptr)
}


// swap function with pointers


func  swap(m1, m2 *int)  {
	// *int we using because both variables are on int type
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp	
}

func main()  {
	m1, m2 := 2, 3
	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)
	
}