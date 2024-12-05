package knapsack

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	Id, Value, Weight int
	IsSelected        bool
}

func MakeItems(numItems, minValue, maxValue, minWeight, maxWeight int) []Item {
	items := make([]Item, numItems)
	for i := range items {
		items[i] = Item{
			Id:         i,
			Value:      minValue + rand.Intn(maxValue-minValue+1),
			Weight:     minWeight + rand.Intn(maxWeight-minWeight+1),
			IsSelected: false,
		}
	}
	return items
}

func CopyItems(items []Item) []Item {
	itemsCopy := make([]Item, len(items))
	copy(itemsCopy, items)
	return itemsCopy
}

func SumValues(items []Item, addAll bool) int {
	sum := 0
	for _, item := range items {
		if addAll || item.IsSelected {
			sum += item.Value
		}
	}
	return sum
}

func SumWeights(items []Item, addAll bool) int {
	sum := 0
	for _, item := range items {
		if addAll || item.IsSelected {
			sum += item.Weight
		}
	}
	return sum
}

func SolutionValue(items []Item, allowedWeight int) int {
	// If the solution's total weight > allowedWeight,
	// return 0 so we won't use this solution.
	if SumWeights(items, false) > allowedWeight {
		return -1
	}
	return SumValues(items, false)
}

func PrintSolution(items []Item) {
	numPrinted := 0
	for i, item := range items {
		if item.IsSelected {
			fmt.Printf("%d(%d, %d) ", i, item.Value, item.Weight)
		}
		numPrinted += 1
		if numPrinted > 100 {
			fmt.Println("...")
			return
		}
	}
	fmt.Println()
}

// Run the algorithm. Display the elapsed time and solution.
func RunAlgorithm(alg func([]Item, int) ([]Item, int, int), items []Item, allowedWeight int) {
	// Copy the items so the run isn't influenced by a previous run.
	testItems := CopyItems(items)

	start := time.Now()

	// Run the algorithm.
	solution, totalValue, functionCalls := alg(testItems, allowedWeight)

	elapsed := time.Since(start)

	fmt.Printf("Elapsed: %f\n", elapsed.Seconds())
	PrintSolution(solution)
	fmt.Printf("Value: %d, Weight: %d, Calls: %d\n",
		totalValue, SumWeights(solution, false), functionCalls)
	fmt.Println()
}
