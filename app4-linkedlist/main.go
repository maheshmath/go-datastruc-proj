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
	fmt.Println("-------------------------")

	linkedList.PrintList()
	fmt.Println("-------------------------")

	linkedList.DeleteNode()
	linkedList.PrintList()
	fmt.Println("-------------------------")

	node := linkedList.SearchNode("Vinanti")
	fmt.Println(node.GetName(), node.GetAge())
	fmt.Println("-------------------------")

	linkedList.DisplayReverseList()
	fmt.Println("-------------------------")

	linkedList.ReverseListUpdated()
	linkedList.PrintList()
	fmt.Println("-------------------------")

	linkedList.InsertAtPosition(3, "Rudra", 7)
	linkedList.PrintList()
	fmt.Println("-------------------------")
}
