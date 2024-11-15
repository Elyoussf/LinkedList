package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func (head *Node) display() {
	fmt.Println("Listing the elements in the linked list:")
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func (head *Node) addAtStart(newdata int) *Node {
	newhead := &Node{data: newdata, next: head}
	return newhead
}

func remove(head *Node, toRemove int) *Node {
	// If head node itself holds the value to be removed
	if head != nil && head.data == toRemove {
		return head.next
	}

	// Search for the node to be deleted, keep track of the previous node
	current := head
	var prev *Node
	for current != nil && current.data != toRemove {
		prev = current
		current = current.next
	}

	// If the node was not found
	if current == nil {
		return head
	}

	// Unlink the node to be deleted
	prev.next = current.next
	return head
}

func main() {
	var head *Node
	head = head.addAtStart(3)
	head = head.addAtStart(2)
	head = head.addAtStart(1)

	head.display()

	fmt.Println("Removing element 2")
	head = remove(head, 2)
	head.display()
}
