package knapsack

func DynamicSearch(items []Item, allowedWeight int) ([]Item, int, int) {
	// Create a 2D slice to store the results of the subproblems.
	// The first index represents the number of items, and the second index represents the allowed weight.
	// The value at each index is the maximum value that can be achieved with the given number of items and allowed weight.
	// The slice is initialized with zeros.
	results := make([][]int, len(items)+1)
	for i := range results {
		results[i] = make([]int, allowedWeight+1)
	}

	// Initialize the first row and column with zeros
	// or with items[0].value when weight can fit.
	// note that the slice is prefilled with zeros.
	for w := 0; w <= allowedWeight; w++ {
		if items[0].Weight <= w {
			results[0][w] = items[0].Value
		}
	}

	// Fill the results slice.
	for i := 1; i <= len(items); i++ {
		for w := 1; w <= allowedWeight; w++ {
			// If the current item's weight is greater than the allowed weight, then the value is the same as the previous item.
			if items[i-1].Weight > w {
				results[i][w] = results[i-1][w]
			} else {
				// Otherwise, the value is the maximum of the previous item and the value of the current item plus the value of the remaining weight.
				results[i][w] = max(results[i-1][w], results[i-1][w-items[i-1].Weight]+items[i-1].Value)
			}
		}
	}

	// Backtrack to find the items that were selected.
	selectedItems := make([]Item, 0)
	i := len(items)
	w := allowedWeight
	for i > 0 && w > 0 {
		if results[i][w] != results[i-1][w] {
			items[i-1].IsSelected = true
			selectedItems = append(selectedItems, items[i-1])
			w -= items[i-1].Weight
		}
		i--
	}

	return selectedItems, results[len(items)][allowedWeight], allowedWeight
}
