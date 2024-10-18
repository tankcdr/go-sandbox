package main

import (
	"fmt"
	"time"

	"github.com/tankcdr/common"
	"github.com/tankcdr/sort"
)

func timeBubbleSort(slice []int) ([]int, time.Duration) {
	startTime := time.Now()

	sort.BubbleSort(slice)

	return slice, time.Since(startTime)
}

func timeOptimizedBubbleSort(slice []int) ([]int, time.Duration) {
	startTime := time.Now()

	sort.OptimizedBubbleSort(slice)

	return slice, time.Since(startTime)
}

func timeCocktailShakerSort(slice []int) ([]int, time.Duration) {
	startTime := time.Now()

	sort.CocktailShakerSort(slice)

	return slice, time.Since(startTime)
}

func timeQuickSort(slice []int) ([]int, time.Duration) {
	startTime := time.Now()

	sort.QuickSort(slice)

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
	_, bubbleSortElapsedTime := timeBubbleSort(append([]int(nil), slice...))
	fmt.Printf("Bubblesort: %v\n", bubbleSortElapsedTime)
	//common.PrintSlice(bubbleSortedSlice, 40)

	_, optBubbleSortElapsedTime := timeOptimizedBubbleSort(append([]int(nil), slice...))
	fmt.Printf("Optimized Bubblesort: %v\n", optBubbleSortElapsedTime)
	//common.PrintSlice(optBubbleSortedSlice, 40)

	_, optCocktailShakerSortElapsedTime := timeCocktailShakerSort(append([]int(nil), slice...))
	fmt.Printf("Cocktail Shaker Sort: %v\n", optCocktailShakerSortElapsedTime)

	_, optQuickSortElapsedTime := timeQuickSort(append([]int(nil), slice...))
	fmt.Printf("Quick Sort: %v\n", optQuickSortElapsedTime)

	// Verify that it's sorted.
	common.CheckSorted(slice)
}
