package common

import (
	"fmt"
	"math/rand"
)

// MakeRandomIntSlice creates a slice of random integers of length n
func MakeRandomIntSlice(n int, max int) []int {
	intSlice := make([]int, n)
	for i := 0; i < n; i++ {
		intSlice[i] = rand.Intn(max)
	}
	return intSlice
}

func PrintSlice(intSlice []int, numItems int) (int, error) {
	if numItems > len(intSlice) {
		numItems = len(intSlice)
	}
	return fmt.Println(intSlice[:numItems])
}

func CheckSorted(slice []int) error {
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return fmt.Errorf("slice is not sorted at index %d", i)
		}
	}
	return nil
}
