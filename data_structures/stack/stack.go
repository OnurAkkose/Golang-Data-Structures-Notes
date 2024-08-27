package main

import "fmt"

// Stack is a struct that holds a slice of integers to represent the stack.
type Stack struct {
	items []int
}

// Push method adds an element to the top of the stack.
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop method removes and returns the top element of the stack.
// If the stack is empty, it prints a message and returns -1.
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("stack is empty")
		return -1
	}

	lastIndex := len(s.items) - 1    // Gets the index of the last item in the slice, which is the top of the stack.
	poppedItem := s.items[lastIndex] // Retrieves the item at the top of the stack.
	s.items = s.items[:lastIndex]    // Removes the top item by slicing off the last element.
	return poppedItem                // Returns the popped item.
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	stack.Pop()

	fmt.Println("Items:")
	for _, i := range stack.items {
		fmt.Println(i)
	}
}
