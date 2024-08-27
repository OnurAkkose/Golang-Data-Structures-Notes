package main

import "fmt"

// Node represents a single node in the binary search tree.
type Node struct {
	value int   // The value stored in the node.
	left  *Node // Pointer to the left child node.
	right *Node // Pointer to the right child node.
}

// NewNode function creates a new node with the given value.
func NewNode(value int) *Node {
	return &Node{value: value}
}

// insert method inserts a new value into the binary search tree.
func (n *Node) insert(value int) {
	if value < n.value { // If the value is less than the current node's value, insert it into the left subtree.
		if n.left == nil { // If the left child is nil, create a new node there.
			n.left = NewNode(value)
		} else {
			n.left.insert(value) // Recursively call insert on the left child.
		}
	} else if value > n.value { // If the value is greater than the current node's value, insert it into the right subtree.
		if n.right == nil { // If the right child is nil, create a new node there.
			n.right = NewNode(value)
		} else {
			n.right.insert(value) // Recursively call insert on the right child.
		}
	}
}

// search method searches for a value in the binary search tree.
// It returns true if the value is found, otherwise false.
func (n *Node) search(value int) bool {
	if n == nil {
		return false
	}

	if value < n.value { // If the value is less than the current node's value, search in the left subtree.
		return n.left.search(value)
	} else if value > n.value { // If the value is greater than the current node's value, search in the right subtree.
		return n.right.search(value)
	}
	return true // If the value matches the current node's value, return true.
}

// inOrderTraversal method performs an in-order traversal of the binary search tree.
// It prints the values of the nodes in ascending order.
func (n *Node) inOrderTraversel() {
	if n != nil {
		n.left.inOrderTraversel() // Recursively traverse the left subtree.
		fmt.Printf("%d ", n.value)
		n.right.inOrderTraversel() // Recursively traverse the right subtree.
	}
}

func main() {
	root := NewNode(7)

	root.insert(1)
	root.insert(55)
	root.insert(3)
	root.insert(21)
	root.insert(12)

	fmt.Println("In-order traversal: ")
	root.inOrderTraversel()
	fmt.Println()

	fmt.Println("Searching for 12: ", root.search(12)) // Should return true.
	fmt.Println("Searching for 15: ", root.search(15)) // Should return false.
}
