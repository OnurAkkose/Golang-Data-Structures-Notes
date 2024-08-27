package main

import "fmt"

// Node represents an element in the linked list.
// It holds a key, a value, and pointers to the next and previous nodes.
type Node struct {
	key   string      // Unique identifier or key for the node (optional, not used in this example)
	value interface{} // The data stored in the node (in this case, a string)
	next  *Node       // Pointer to the next node in the list
	prev  *Node       // Pointer to the previous node in the list (not used in this example)
}

// LinkedList represents a doubly linked list.
// It maintains a pointer to the head (first node) of the list.
type LinkedList struct {
	head *Node // Pointer to the first node in the linked list
}

// append adds a new node with the specified value to the end of the linked list.
func (list *LinkedList) append(value string) {
	// Create a new node with the given value
	node := &Node{value: value}

	// If the list is empty (head is nil), set the new node as the head
	if list.head == nil {
		list.head = node
		return
	}

	// Traverse the list to find the last node
	current := list.head
	for current.next != nil {
		current = current.next
	}

	// Set the new node as the next node of the last node
	current.next = node
}

// prepend adds a new node with the specified value to the beginning of the linked list.
func (list *LinkedList) prepend(value string) {
	// Create a new node with the given value
	node := &Node{value: value}

	// Set the new node's next pointer to the current head of the list
	node.next = list.head

	// Update the head to the new node
	list.head = node
}

// delete removes the first node with the specified value from the linked list.
func (list *LinkedList) delete(value string) {
	// If the list is empty, there's nothing to delete
	if list.head == nil {
		return
	}

	// If the head node contains the value, update the head to the next node
	if list.head.value == value {
		list.head = list.head.next
		return
	}

	// Traverse the list to find the node with the matching value
	current := list.head
	for current.next != nil {
		if current.next.value == value {
			// If found, bypass the node by updating the next pointer of the previous node
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

// printListNodes prints all the values in the linked list in order.
func (list *LinkedList) printListNodes() {
	fmt.Print("Values: ")

	// Start from the head and traverse the list, printing each node's value
	current := list.head
	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}

	// Indicate the end of the list
	fmt.Println("nil")
}

func main() {
	linkedList := LinkedList{}
	linkedList.append("a")
	linkedList.append("b")
	linkedList.append("c")
	linkedList.append("d")

	linkedList.delete("b")

	linkedList.printListNodes()
}
