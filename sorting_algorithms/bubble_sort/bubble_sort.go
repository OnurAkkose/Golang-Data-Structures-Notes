package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Printf("at %d iteration, sorted as %v\n", i, arr)
	}
}

func main() {
	arr := []int{63, 34, 25, 12, 22, 11, 91}

	bubbleSort(arr)

	fmt.Println("Sorted array:", arr)
}
