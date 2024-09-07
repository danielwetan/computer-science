package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func reverseBetween(head *Node, m int, n int) *Node {
	if head == nil || m == n {
		return head
	}

	var previous *Node
	current := head

	// Move to the start of the sublist (position m)
	for i := 1; i < m; i++ {
		previous = current
		current = current.next
	}
	// previous -> 1 2 3 4 5
	// current  -> 2 3 4 5

	// Save the positions for connection later
	lastNodeOfFirstPart := previous // Points to the node at index 'm-1'
	lastNodeOfSubList := current    // After reversal, 'current' becomes the last node of the sublist
	var next *Node                  // Temporarily stores the next node

	// Reverse the sublist between 'm' and 'n'
	for i := 0; i < n-m+1; i++ {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}

	// Connect the first part to the reversed sublist
	if lastNodeOfFirstPart != nil {
		lastNodeOfFirstPart.next = previous
	} else {
		head = previous // If m = 1, the new head is the start of the reversed sublist
	}

	lastNodeOfSubList.next = current

	return head
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
	// Example: Reverse sublist from position 2 to 4
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	// printLinkedList(head) // Output: 1 2 3 4 5

	head = reverseBetween(head, 2, 4)
	printLinkedList(head) // Output: 1 4 3 2 5
}
