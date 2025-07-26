package main

import (
	"fmt"
	"sorting/sorting"
)

func main() {
	var n int
	fmt.Print("Enter size of array: ")
	fmt.Scan(&n)

	original := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter element %d: ", i)
		fmt.Scan(&original[i])
	}

	fmt.Print("Unsorted Array: ")
	for _, val := range original {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	insertionInput := make([]int, n)
	copy(insertionInput, original)
	insertionSorted := sorting.InsertionSort(insertionInput)
	fmt.Print("Sorted Array using Insertion Sort: ")
	for _, val := range insertionSorted {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	mergeInput := make([]int, n)
	copy(mergeInput, original)
	mergeSorted := sorting.MergeSort(mergeInput)
	fmt.Print("Sorted Array using Merge Sort: ")
	for _, val := range mergeSorted {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	selectionInput := make([]int, n)
	copy(selectionInput, original)
	selectionSorted := sorting.SelectionSort(selectionInput)
	fmt.Print("Sorted Array using Selection Sort: ")
	for _, val := range selectionSorted {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	bubbleInput := make([]int, n)
	copy(bubbleInput, original)
	bubbleSorted := sorting.BubbleSort(bubbleInput)
	fmt.Print("Sorted Array using Bubble Sort: ")
	for _, val := range bubbleSorted {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	quickInput := make([]int, n)
	copy(quickInput, original)
	quickSorted := sorting.QuickSort(quickInput)
	fmt.Print("Sorted Array using Quick Sort: ")
	for _, val := range quickSorted {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
