package main

import (
	"fmt"
	"go-datastruc-pro/app4-linkedlist/entities"
)

func main() {

	var linkedList = entities.LinkedList{}
	linkedList.AddNode("Mahesh", 43)
	linkedList.AddNode("Vinanti", 38)
	linkedList.AddNode("Aarush", 13)
	linkedList.AddNode("Rudra", 7)

	fmt.Println(linkedList.Head.GetName())

	linkedList.PrintList()
	linkedList.DeleteNode()

	linkedList.PrintList()
}
