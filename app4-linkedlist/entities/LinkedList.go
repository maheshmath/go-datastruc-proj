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

func (l *LinkedList) DisplayReverseList() {
	currentNode := l.Head
	if currentNode == nil {
		return
	}

	l.traverse(currentNode)

}

func (l *LinkedList) traverse(node *Node) {
	if node == nil {
		return
	}
	l.traverse(node.next)
	fmt.Println(node.name, node.age)
}

func (l *LinkedList) ReverseList() {
	node := l.Head
	if node == nil {
		return
	}

	l.reverse(node)
}

func (l *LinkedList) reverse(node *Node) {
	if node.next == nil {
		l.Head = node.next
		return
	}
	l.reverse(node.next)
	node.next.next = node
	node.next = nil
}

func (l *LinkedList) InsertAtPosition(position int, name string, age int) {
	newNode := Node{name, age, nil}
	if l.Head == nil {
		l.Head = &newNode
		return
	}

	currentNode := l.Head
	for index := 1; index < position-1; index++ {
		currentNode = currentNode.next
	}

	currentNode.next = &newNode
	newNode.next = currentNode.next
}
