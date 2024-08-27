package main

import "fmt"

func createArray() []int {
	return []int{1, 2, 3, 4, 5, 6, 7}
}

func iterateArray(arr []int) {
	for index, item := range arr {
		fmt.Println()
		fmt.Printf("index: %d and value: %d", index, item)
	}
}

func updateArray(arr []int) []int {
	arr[1] = 10                    // Updates the value at index 1 to 10.
	newArray := []int{10, 11, 12}  // Creates a new slice with additional values.
	arr = append(arr, newArray...) // Appends the new slice to the original array.
	return arr
}

func removeItemByIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...) // Removes the element at the specified index and returns the modified array.
}

func main() {
	myArray := createArray()
	fmt.Println("Created array => ", myArray)

	iterateArray(myArray)

	fmt.Println()
	updatedArray := updateArray(myArray)
	fmt.Println("Updated array => ", updatedArray)

	updatedArray = removeItemByIndex(updatedArray, 2)
	fmt.Println("Array after removing the 2nd index => ", updatedArray)
}
