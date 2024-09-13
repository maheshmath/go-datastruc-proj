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

func (node *Node) GetAge() int {
	return node.age
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

func (l *LinkedList) SearchNode(name string) *Node {
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.name == name {
			return currentNode
		}
		currentNode = currentNode.next
	}

	return nil
}

func (l *LinkedList) ReverseList() {
	currentNode := l.Head
	if currentNode == nil {
		return
	}

	l.reverse(currentNode)

}

func (l *LinkedList) reverse(node *Node) {
	if node == nil {
		return
	}

	l.reverse(node.next)
	fmt.Println(node.name, node.age)
}
