package main

import "fmt"

type SinglyLinkedList struct {
	head   *Node
	length int
}

func (s *SinglyLinkedList) add(node *Node) {
	if s.head == nil {
		s.head = node
		return
	}

	currentNode := s.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = node

	s.length++
}

func (s *SinglyLinkedList) toString() {
	currentNode := s.head
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.value)
		currentNode = currentNode.next
	}
}

type Node struct {
	value int
	next  *Node
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func main() {

	singlyLinkedList := NewSinglyLinkedList()
	singlyLinkedList.add(&Node{value: 1})
	singlyLinkedList.add(&Node{value: 2})
	singlyLinkedList.add(&Node{value: 3})

	singlyLinkedList.toString()
}
