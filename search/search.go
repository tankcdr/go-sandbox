package search

func LinearSearch(A []int, target int) (index int, numTests int) {
	tests := 0

	for i, v := range A {
		tests++
		if v == target {
			return i, tests
		}
	}
	return -1, tests
}

func BinarySearch(A []int, target int) (index int, numTests int) {
	tests := 0

	lo := 0
	hi := len(A) - 1

	for lo <= hi {
		tests++
		mid := lo + (hi-lo)/2

		if A[mid] == target {
			return mid, tests
		} else if A[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return -1, tests
}
