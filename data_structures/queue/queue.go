package main

import "fmt"

// Queue is a struct that holds a slice of integers to represent the queue.
type Queue struct {
	items []int
}

// Enqueue method adds an element to the end of the queue.
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue method removes and returns the front element of the queue.
// If the queue is empty, it prints a message and returns -1.
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	dequeuedItem := q.items[0] // Retrieves the item at the front of the queue.
	q.items = q.items[1:]      // Removes the front item by re-slicing the slice to exclude the first element.
	return dequeuedItem        // Returns the dequeued item.
}

func main() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	queue.Dequeue()

	fmt.Println("Items:")
	for _, i := range queue.items {
		fmt.Println(i)
	}
}
