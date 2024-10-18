package search_test

import (
	"testing"

	"github.com/tankcdr/common"
	"github.com/tankcdr/search"
	"github.com/tankcdr/sort"
)

func TestSearch_LinearSearch(t *testing.T) {
	arr := common.MakeRandomIntSlice(1000, 1000)
	target := 1001
	arr[500] = target

	index, numTests := search.LinearSearch(arr, target)

	if arr[index] != target {
		t.Error("Search error: target is not at index")
	}

	if numTests != 501 {
		t.Errorf("Search error: incorrect number of tests %d", numTests)
	}
}

func TestSearch_LinearSearch_NotFound(t *testing.T) {
	arr := common.MakeRandomIntSlice(1000, 1000)
	target := 1001

	index, numTests := search.LinearSearch(arr, target)

	if index != -1 {
		t.Error("Search error: target found when it shouldn't be")
	}

	if numTests != 1000 {
		t.Errorf("Search error: incorrect number of tests %d", numTests)
	}
}

func TestSearch_BinarySearch(t *testing.T) {
	arr := common.MakeRandomIntSlice(500, 10000)
	sort.QuickSort(arr)

	targetIndex := 31
	target := arr[targetIndex]

	index, _ := search.BinarySearch(arr, target)

	if index != targetIndex {
		t.Errorf("Search error: target is not at index, %d", index)
	}
}
