package main

import (
	"fmt"
)

/**
go has a garbage collector which means you don't have to
deal with memory management
But you may have other kind of processing which uses your resources
like file, socket, virtual machine etc
You would like to make sure these resource are also close
when you programe is closed
to make sure your resources is close use defer
*/
func cleanup(name string) {
	fmt.Println("cleaning up %s\n", name)
}

func main() {
	defer cleanup("A")
	defer cleanup("B")
	fmt.Println("worker")

	/**
	in output first worker is print then defer function call
	defer is call when some error comes in your code
	defer call in reverse order
	*/
}