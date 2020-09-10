// used for go comments
/**
used for multi line comments
*/
// everything in go is in form of package
package main //main package make your code compile and execute
// e.g if we have a 3-4 files containing package main imported then compiler combine all files in a main file which is called main.go
// its helping in modularity

import (
	"fmt" //fmt is for formating printed line and logging
)

func HelloWorld()  {
	fmt.Println("i am a hello world")
}
// Every file has a main function which is a entry point of every file
func main() {
	fmt.Printf("hello, world\n")
}
// to run this file go to terminal and run "go run filename"
// to build this file we have to run "go build filename"
// go test => have build in test cases
// go get => to install third party packages
