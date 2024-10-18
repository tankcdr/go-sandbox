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

func LomutoPartition(A []int) (pivot int, index int) {
	lo := 0
	hi := len(A) - 1
	pivot = A[hi]

	for j := lo; j < hi; j++ {
		if A[j] < pivot {
			A[lo], A[j] = A[j], A[lo]
			lo++
		}
	}

	A[lo], A[hi] = A[hi], A[lo]
	return pivot, lo
}

func QuickSort(A []int, partitions ...func([]int) (int, int)) {
	var partition func([]int) (int, int)

	if len(partitions) > 0 && partitions[0] != nil {
		partition = partitions[0]
	} else {
		partition = LomutoPartition
	}

	if len(A) <= 1 {
		return
	}

	_, p := partition(A)
	QuickSort(A[:p], partition)
	QuickSort(A[p+1:], partition)
}

func CountingSort(A []int, exclusiveMax int) []int {
	B := make([]int, len(A))
	C := make([]int, exclusiveMax)

	for _, v := range A {
		C[v]++
	}

	for index, _ := range C {
		if index > 0 {
			C[index] += C[index-1]
		}
	}

	for i := len(A) - 1; i >= 0; i-- {
		v := A[i]
		C[v]--
		B[C[v]] = v
	}

	return B
}
