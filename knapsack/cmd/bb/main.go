package main

import (
	"fmt"
	"knapsack"
)

const numItems = 32

const minValue = 1
const maxValue = 10
const minWeight = 4
const maxWeight = 10

var allowedWeight int

func main() {
	//items := makeTestItems()
	items := knapsack.MakeItems(numItems, minValue, maxValue, minWeight, maxWeight)
	allowedWeight = knapsack.SumWeights(items, true) / 2

	// Display basic parameters.
	fmt.Println("*** Parameters ***")
	fmt.Printf("# items: %d\n", numItems)
	fmt.Printf("Total value: %d\n", knapsack.SumValues(items, true))
	fmt.Printf("Total weight: %d\n", knapsack.SumWeights(items, true))
	fmt.Printf("Allowed weight: %d\n", allowedWeight)
	fmt.Println()

	// Exhaustive search
	if numItems > 65 {
		fmt.Println("Too many items for exhaustive search")
	} else {
		fmt.Println("*** Branch and Bound Search ***")
		knapsack.RunAlgorithm(knapsack.BranchAndBoundSearch, items, allowedWeight)
	}
}
