package main

import "fmt"

type Node struct {
	value int

	next *Node
	prev *Node

	head *Node
	tail *Node
}

func (n *Node) add(value int) {
	node := Node{
		value: value,
		head:  n.head,
	}

	n.next = &node
	node.prev = n

	current := n.head
	for current != nil {
		current.tail = &node
		current = current.next
	}
}

func (n *Node) displayHead() {
	fmt.Println(n.head.value)
}

func (n *Node) displayTail() {
	fmt.Println(n.tail.value)
}

func (n *Node) displayList() {
	current := n.head
	for current != nil {
		fmt.Print(current.value, "\t")
		current = current.next
	}

	fmt.Println()
}

func (n *Node) reverse() (list *Node) {
	current := n.tail

	list = NewList(current.value)
	current = current.prev

	for current != nil {
		list.add(current.value)
		list = list.next

		current = current.prev
	}

	return
}

func NewList(value int) (list *Node) {
	list = &Node{
		value: value,
		next:  nil,
	}

	list.tail = list
	list.head = list

	return
}

func task6() {
	list := NewList(10)

	list.add(50)
	list = list.next

	list.add(60)
	list = list.next

	list.add(70)
	list = list.next

	list.displayHead()
	list.displayTail()
	list.displayList()

	list = list.reverse()

	list.displayHead()
	list.displayTail()
	list.displayList()
}
