package knapsack

import "sort"

type BlockableItem struct {
	Item
	BlockedBy   int
	BlockedList []int
}

func RodsTechniqueSearch(items []Item, allowedWeight int) ([]Item, int, int) {

	//wrap and setup
	blockableItems := make([]BlockableItem, len(items))
	for i, item := range items {
		blockableItems[i] = BlockableItem{
			Item:        item,
			BlockedBy:   -1,
			BlockedList: nil,
		}
	}
	makeBlockLists(blockableItems)

	//find solution
	solutionItems, value, calls := doRodsTechniqueSearch(blockableItems, allowedWeight, 0, 0, 0, 0, sumValues(blockableItems, true))

	//unwrap
	solution := make([]Item, len(solutionItems))
	for i, item := range solutionItems {
		solution[i] = item.Item
	}

	return solution, value, calls
}

// sorts blackable items by length of blocked list
// this is a simple heuristic to reduce the number of calls
// by moving longer lists to the front of the search
func RodsTechniqueSortedSearch(items []Item, allowedWeight int) ([]Item, int, int) {

	//wrap and setup
	blockableItems := make([]BlockableItem, len(items))
	for i, item := range items {
		blockableItems[i] = BlockableItem{
			Item:        item,
			BlockedBy:   -1,
			BlockedList: nil,
		}
	}
	makeBlockLists(blockableItems)

	//sort by length of blocked list
	sort.Slice(blockableItems, func(i, j int) bool {
		return len(blockableItems[i].BlockedList) > len(blockableItems[j].BlockedList)
	})

	//fix items
	// After sorting, update the Ids
	for idx := range blockableItems {
		blockableItems[idx].Id = idx
	}
	//rerun makeBlockLists
	makeBlockLists(blockableItems)

	//find solution
	solutionItems, value, calls := doRodsTechniqueSearch(blockableItems, allowedWeight, 0, 0, 0, 0, sumValues(blockableItems, true))

	//unwrap
	solution := make([]Item, len(solutionItems))
	for i, item := range solutionItems {
		solution[i] = item.Item
	}

	return solution, value, calls
}

func doRodsTechniqueSearch(items []BlockableItem, allowedWeight, nextIndex, bestValue, currentValue, currentWeight, remainingValue int) ([]BlockableItem, int, int) {

	if nextIndex >= len(items) {
		return copyItems(items), currentValue, 1
	}

	// Pruning condition
	if currentValue+remainingValue <= bestValue {
		return nil, currentValue, 1
	}

	// Try adding the next item.
	var test1Solution []BlockableItem
	test1Solution = nil
	test1Value := 0
	test1Calls := 1

	// Branch where the current item is selected
	if currentWeight+items[nextIndex].Weight <= allowedWeight && items[nextIndex].BlockedBy == -1 {
		items[nextIndex].IsSelected = true

		test1Solution, test1Value, test1Calls = doRodsTechniqueSearch(
			items, allowedWeight, nextIndex+1, bestValue,
			currentValue+items[nextIndex].Value,
			currentWeight+items[nextIndex].Weight,
			remainingValue-items[nextIndex].Value)

		if test1Value > bestValue {
			bestValue = test1Value
		}
	}

	// Branch where the current item is not selected
	blockItems(items[nextIndex], items)
	items[nextIndex].IsSelected = false
	test2Solution, test2Value, test2Calls := doRodsTechniqueSearch(
		items, allowedWeight, nextIndex+1, bestValue, currentValue,
		currentWeight, remainingValue-items[nextIndex].Value)
	unblockItems(items[nextIndex], items)

	// Return the solution that is better.
	if test1Value >= test2Value {
		return test1Solution, test1Value, test1Calls + test2Calls + 1
	} else {
		return test2Solution, test2Value, test1Calls + test2Calls + 1
	}
}

func copyItems(items []BlockableItem) []BlockableItem {
	itemsCopy := make([]BlockableItem, len(items))
	for i := range items {
		itemsCopy[i] = items[i]
		// If BlockedList is a slice, copy it to prevent shared references
		if items[i].BlockedList != nil {
			itemsCopy[i].BlockedList = make([]int, len(items[i].BlockedList))
			copy(itemsCopy[i].BlockedList, items[i].BlockedList)
		}
	}
	return itemsCopy
}

func solutionValue(items []BlockableItem, allowedWeight int) int {
	// If the solution's total weight > allowedWeight,
	// return 0 so we won't use this solution.
	if sumWeights(items, false) > allowedWeight {
		return -1
	}
	return sumValues(items, false)
}

func sumValues(items []BlockableItem, addAll bool) int {
	sum := 0
	for _, item := range items {
		if addAll || item.IsSelected {
			sum += item.Value
		}
	}
	return sum
}

func sumWeights(items []BlockableItem, addAll bool) int {
	sum := 0
	for _, item := range items {
		if addAll || item.IsSelected {
			sum += item.Weight
		}
	}
	return sum
}

func makeBlockLists(items []BlockableItem) {
	for i := range items {
		items[i].BlockedList = nil
		for j, otherItem := range items {
			if i != j {
				if items[i].Weight <= otherItem.Weight && items[i].Value >= otherItem.Value {
					items[i].BlockedList = append(items[i].BlockedList, j)
				}
			}
		}
	}
}

func blockItems(source BlockableItem, items []BlockableItem) {
	for _, index := range source.BlockedList {
		if items[index].BlockedBy == -1 {
			items[index].BlockedBy = source.Id
		}

	}
}

func unblockItems(source BlockableItem, items []BlockableItem) {
	for _, index := range source.BlockedList {
		if items[index].BlockedBy == source.Id {
			items[index].BlockedBy = -1
		}
	}
}
