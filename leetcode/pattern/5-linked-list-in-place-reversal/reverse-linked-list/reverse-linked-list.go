package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

func reverseLinkedList(head *Node) *Node {
	var (
		prev    *Node
		current *Node
		next    *Node
	)

	current = head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current

		current = next
	}

	return prev
}

func createLinkedList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}

	head := &Node{value: arr[0]}
	current := head
	for _, v := range arr[1:] {
		current.next = &Node{value: v}
		current = current.next
	}
	return head
}

// Helper function to print the linked list
func printLinkedList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.value)
		head = head.next
	}
	fmt.Println()
}

func main() {
	head := createLinkedList([]int{1, 2, 3, 4})
	head = reverseLinkedList(head)
	printLinkedList(head) // Output: 4 3 2 1
}
