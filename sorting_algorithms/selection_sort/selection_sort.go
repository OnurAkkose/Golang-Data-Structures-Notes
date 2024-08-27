package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Find the minimum element in the unsorted portion of the array
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap the found minimum element with the first element of the unsorted portion
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Unsorted array:", arr)

	selectionSort(arr)
	fmt.Println("Sorted array:", arr)
}
