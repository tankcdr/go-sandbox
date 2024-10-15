package main

import (
	"fmt"
	"time"

	"github.com/tankcdr/common"
	"github.com/tankcdr/sort"
)

func timeBubblesort(slice []int) ([]int, time.Duration) {
	startTime := time.Now()

	sort.Bubblesort(slice)

	return slice, time.Since(startTime)
}

func timeOptimizedBubblesort(slice []int) ([]int, time.Duration) {
	startTime := time.Now()

	sort.OptimizedBubbleSort(slice)

	return slice, time.Since(startTime)
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := common.MakeRandomIntSlice(numItems, max)
	//common.PrintSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	_, bubbleSortElapsedTime := timeBubblesort(append([]int(nil), slice...))
	fmt.Printf("Bubblesort: %v\n", bubbleSortElapsedTime)
	//common.PrintSlice(bubbleSortedSlice, 40)

	_, optBubbleSortElapsedTime := timeOptimizedBubblesort(append([]int(nil), slice...))
	fmt.Printf("Optimized Bubblesort: %v\n", optBubbleSortElapsedTime)
	//common.PrintSlice(optBubbleSortedSlice, 40)

	// Verify that it's sorted.
	common.CheckSorted(slice)
}
