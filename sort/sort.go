package sort

// BubbleSort sorts a slice of integers using the bubblesort algorithm.
func BubbleSort(a []int) {
	for itemCount := len(a); ; {
		swapped := false
		for i := 1; i < itemCount; i++ {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func OptimizedBubbleSort(a []int) {
	for itemCount := len(a); ; {
		newItem := 0
		for i := 1; i < itemCount; i++ {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				newItem = i
			}
		}
		if newItem == 0 {
			break
		}
		itemCount = newItem
	}
}

func CocktailShakerSort(a []int) {
	for itemCount := len(a); ; {
		newItem := 0
		for i := 1; i < itemCount; i++ {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				newItem = i
			}
		}
		if newItem == 0 {
			break
		}
		itemCount = newItem

		for i := itemCount - 1; i > 0; i-- {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
			}
		}
	}
}
