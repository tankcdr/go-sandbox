package knapsack

func BranchAndBoundSearch(items []Item, allowedWeight int) ([]Item, int, int) {
	//bestValue, currentValue, currentWeight := 0, 0, 0
	//remainingValue := SumValues(items, true)
	return doBranchAndBoundSearch(items, allowedWeight, 0, 0, 0, 0, SumValues(items, true))
}

func doBranchAndBoundSearch(items []Item, allowedWeight, nextIndex, bestValue, currentValue, currentWeight, remainingValue int) ([]Item, int, int) {

	// Initialize variables for the two branches.
	// test1 is the branch where the current item is selected.
	var test1Solution []Item
	test1Value, test1Calls := 0, 0
	// test2 is the branch where the current item is not selected.
	var test2Solution []Item
	test2Value, test2Calls := 0, 0

	if nextIndex >= len(items) {
		return CopyItems(items), SolutionValue(items, allowedWeight), 1
	}

	if currentValue+remainingValue <= bestValue {
		return nil, 0, 1
	}

	if currentWeight+items[nextIndex].Weight <= allowedWeight {
		items[nextIndex].IsSelected = true
		test1Solution, test1Value, test1Calls = doBranchAndBoundSearch(items, allowedWeight, nextIndex+1, bestValue, currentValue+items[nextIndex].Value, currentWeight+items[nextIndex].Weight, remainingValue-items[nextIndex].Value)

	} else {
		test1Solution, test1Value, test1Calls = nil, 0, 1
	}

	if currentValue+remainingValue-items[nextIndex].Value > bestValue {
		items[nextIndex].IsSelected = false
		test2Solution, test2Value, test2Calls = doBranchAndBoundSearch(items, allowedWeight, nextIndex+1, bestValue, currentValue, currentWeight, remainingValue-items[nextIndex].Value)

	} else {
		test2Solution, test2Value, test2Calls = nil, 0, 1
	}

	if test1Value > test2Value {
		return test1Solution, test1Value, test1Calls + test2Calls
	}

	return test2Solution, test2Value, test1Calls + test2Calls
}
