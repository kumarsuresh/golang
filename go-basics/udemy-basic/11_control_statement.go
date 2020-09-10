package main

import  (
	"fmt"
)

func main() {
	fmt.Print("Hello world!!")
	//if else, for, switch case, breat continue
	f := false
	flag := &f

	if  flag == nil {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("Value is True")
	} else {
		fmt.Println("Value is False")
	}

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
	// for infinite loop
	// for {
	// 	fmt.Println("skldj")
	// }

	arr := []string{"Kumar", "Suresh", "Ratten"}
	for i, value:= range arr {
		fmt.Println(i)
		fmt.Print(value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "Suresh"
	mymap["age"] = 34

	for k, v := range mymap {
		fmt.Printf("\nKey: %s value is %v\n", k, v)
	}
}
