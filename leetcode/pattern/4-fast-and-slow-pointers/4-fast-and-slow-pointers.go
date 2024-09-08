package main

import "fmt"

type Node struct {
	val  string
	next *Node
}

func hasCycle(head *Node) bool {
	if head == nil || head.next == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next      // move slow pointer by 1 step
		fast = fast.next.next // move fast pointer by 2 step

		if slow == fast {
			return true // cycle detected
		}
	}

	return false
}

func main() {
	a := &Node{val: "a"}
	b := &Node{val: "b"}
	c := &Node{val: "c"}

	a.next = b
	b.next = c

	fmt.Println("hasCycle? ", hasCycle(a))

	x := &Node{val: "x"}
	y := &Node{val: "y"}
	z := &Node{val: "z"}

	x.next = y
	y.next = z
	z.next = y

	fmt.Println("hasCycle? ", hasCycle(x))
}
