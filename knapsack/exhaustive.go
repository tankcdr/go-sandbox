package knapsack

func ExhaustiveSearch(items []Item, allowedWeight int) ([]Item, int, int) {
	return doExhaustiveSearch(items, allowedWeight, 0)
}

func doExhaustiveSearch(items []Item, allowedWeight, nextIndex int) ([]Item, int, int) {
	if nextIndex >= len(items) {
		return CopyItems(items), SolutionValue(items, allowedWeight), 1
	}

	items[nextIndex].IsSelected = true
	leftItems, leftValue, leftCount := doExhaustiveSearch(items, allowedWeight, nextIndex+1)

	items[nextIndex].IsSelected = false
	rightItems, rightValue, rightCount := doExhaustiveSearch(items, allowedWeight, nextIndex+1)

	if leftValue > rightValue {
		return leftItems, leftValue, leftCount + rightCount
	}

	return rightItems, rightValue, leftCount + rightCount
}
