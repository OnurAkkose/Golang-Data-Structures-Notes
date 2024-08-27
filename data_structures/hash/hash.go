package main

import "fmt"

const ArraySize = 10

// HashTable represents the hash table data structure
type HashTable struct {
	array [ArraySize]*bucket // The hash table is an array of pointers to buckets (linked lists).
}

// bucket is a linked list used to handle collisions in the hash table
type bucket struct {
	head *bucketNode // The bucket contains a linked list of bucketNodes.
}

// bucketNode represents a node in the linked list (bucket)
type bucketNode struct {
	key  string      // The key stored in the node.
	next *bucketNode // Pointer to the next node in the linked list.
}

// CreateHashTable initializes a new hash table with empty buckets
func CreateHashTable() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{} // Initialize each bucket as an empty linked list.
	}
	return result
}

// insert adds a new key to the bucket by creating a new node and inserting it at the head
func (b *bucket) insert(key string) {
	node := &bucketNode{key: key} // Create a new bucketNode with the key.
	node.next = b.head            // Set the new node's next to the current head of the bucket.
	b.head = node                 // Update the head of the bucket to the new node.
}

// insert adds a key to the hash table by hashing it to determine the bucket index
func (h *HashTable) insert(key string) {
	index := hash(key)         // Hash the key to find the index.
	h.array[index].insert(key) // Insert the key into the appropriate bucket.
}

// search checks if a key exists in the bucket by iterating through the linked list
func (b *bucket) search(key string) bool {
	currentNode := b.head    // Start from the head of the bucket.
	for currentNode != nil { // Iterate through the linked list.
		if key == currentNode.key { // If the key is found, return true.
			return true
		}
		currentNode = currentNode.next // Move to the next node.
	}
	return false // If the key is not found, return false.
}

// search checks if a key exists in the hash table by hashing it to determine the bucket index
func (h *HashTable) search(key string) bool {
	index := hash(key)                // Hash the key to find the index.
	return h.array[index].search(key) // Search for the key in the appropriate bucket.
}

// delete removes a key from the bucket by adjusting the linked list pointers
func (b *bucket) delete(key string) bool {
	// Check if the head node contains the key
	if b.head != nil && b.head.key == key {
		b.head = b.head.next // Remove the head node.
		return true
	}

	previousNode := b.head // Start from the head of the bucket.
	for previousNode != nil && previousNode.next != nil {
		if previousNode.next.key == key { // If the next node contains the key, delete it.
			previousNode.next = previousNode.next.next // Bypass the node to delete it.
			return true
		}
		previousNode = previousNode.next // Move to the next node.
	}
	return false // If the key is not found, return false.
}

// delete removes a key from the hash table by hashing it to determine the bucket index
func (h *HashTable) delete(key string) bool {
	index := hash(key)                // Hash the key to find the index.
	return h.array[index].delete(key) // Delete the key from the appropriate bucket.
}

// hash is a simple hash function that sums the ASCII values of the key's characters
// and returns an index based on the size of the hash table array
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v) // Sum the ASCII values of the characters in the key.
	}
	return sum % ArraySize // Return the sum modulo the array size to get the index.
}

func main() {
	hashTable := CreateHashTable() // Create a new hash table.
	list := []string{
		"ONUR",
		"AKKÖSE",
		"SOFTWARE",
		"DEMO",
		"DATA STRUCTURES",
	}

	// Insert keys into the hash table
	for _, v := range list {
		hashTable.insert(v)
	}

	fmt.Println(hashTable)

	fmt.Println(hashTable.delete("DEMO"))

	fmt.Println(hashTable)

	fmt.Println(hashTable.search("SOFTWARE"))
	fmt.Println(hashTable.search("test"))
}
