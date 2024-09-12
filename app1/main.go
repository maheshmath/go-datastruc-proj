package main

import (
	"fmt"
	"go-datastruc-pro/app1/entities"
)

func main() {
	person := entities.Person{
		Name: "John",
		Age:  30,
	}
	fmt.Println(person)
}
