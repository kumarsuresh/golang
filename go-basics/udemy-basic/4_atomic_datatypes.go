package main
import (
	"fmt"
)

func main()  {
	var m1 int
	m1 = 2
	fmt.Println(m1)
}


func main()  {
	// to define multiple variable 
	var (
		m1 = 2 // no need to define datatype because we use = sign
		m2 = 3
	)
	fmt.Println(m1 + m2)
}


//mismatch data types
// type casting
func main()  {
	var m1 int64
	var m2 int32
	
	fmt.Println(m1 + int64(m2))
	// ans is zero because all golang int default value is 0
}

// initialize variable in one line

func main()  {
	m1 := 1
	m2 := 2
	fmt.Println(m1 + m2)
}