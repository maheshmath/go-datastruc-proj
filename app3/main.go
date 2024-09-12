package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello, World!")

	if 5 == 4 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var values map[string]string = make(map[string]string)
	values["Name1"] = "Mahesh"
	values["Name2"] = "Rudra"
	fmt.Println(values)

	for key, value := range values {
		fmt.Println(key, value)
	}
}
