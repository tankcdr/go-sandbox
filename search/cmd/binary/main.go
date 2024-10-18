package main

import (
	"fmt"
	"time"

	"github.com/tankcdr/common"
	"github.com/tankcdr/search"
	"github.com/tankcdr/sort"
)

func timeBinarySearch(slice []int, target int) (index int, numTessts int, runTime time.Duration) {
	startTime := time.Now()

	index, numTests := search.BinarySearch(slice, target)

	return index, numTests, time.Since(startTime)
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
	slice = sort.CountingSort(slice, max)
	//common.PrintSlice(slice, 40)
	fmt.Printf("slice: %v\n", slice)

	for ok := true; ok; {
		// Get the target value.
		var target int
		fmt.Printf("Target: ")
		read, _ := fmt.Scanln(&target)

		if read == 0 {
			fmt.Print("No input provided. Exiting...\n")
			break
		}

		fmt.Printf("Searching for %d in slice...\n", target)

		// Search for the target value.
		index, numTests, runTime := timeBinarySearch(slice, target)

		// Display the search results.
		if index == -1 {
			fmt.Printf("Target %d not found in slice\n", target)
		} else {
			fmt.Printf("Target %d found at index %d\n", target, index)
		}
		fmt.Printf("Number of tests: %d, Runtime: %v\n", numTests, runTime)
		fmt.Println()

	}
}
