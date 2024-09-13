package entities

import "fmt"

type Node struct {
	name string
	age  int
	next *Node
}

func (node *Node) GetName() string {
	return node.name
}

type LinkedList struct {
	Head *Node
}

// AddNode adds a new node at the end of the linked list.
func (l *LinkedList) AddNode(name string, age int) {
	tempNode := Node{name, age, nil}

	if l.Head == nil {
		l.Head = &tempNode
	} else {
		currentNode := l.Head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = &tempNode
	}
}

func (l *LinkedList) DeleteNode() bool {
	if l.Head == nil {
		return false
	}

	currentNode := l.Head
	for currentNode.next.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = nil
	return true
}

func (l *LinkedList) PrintList() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Println(currentNode.name, currentNode.age)
		currentNode = currentNode.next
	}
}
