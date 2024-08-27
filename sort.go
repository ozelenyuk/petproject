package main

import (
	"fmt"
)

// Function to perform bubble sort on a slice
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Last i elements are already sorted
		for j := 0; j < n-i-1; j++ {
			// Swap if the element found is greater than the next element
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// Function to print a slice
func printArray(arr []int) {
	for _, value := range arr {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original array:")
	printArray(arr)

	bubbleSort(arr)

	fmt.Println("Sorted array:")
	printArray(arr)
}
