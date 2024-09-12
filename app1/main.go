package main

import (
	"fmt"
	"go-datastruc-pro/app1/entities"
)

func main() {
	var person entities.Person
	fmt.Printf("Name is %s and age is %d", person.GetName(), person.GetAge()) // Name is %s and age is {%d}", person.GetName(), person.GetAge())
}
