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
