package sort_test

import (
	"testing"

	"github.com/tankcdr/common"
	"github.com/tankcdr/sort"
)

func TestSort_Bubblesort(t *testing.T) {
	arr := common.MakeRandomIntSlice(1000, 1000)

	sort.BubbleSort(arr)
	err := common.CheckSorted(arr)

	if err != nil {
		t.Error(err)
	}
}

func TestSort_OptimizedBubbleSort(t *testing.T) {
	arr := common.MakeRandomIntSlice(1000, 1000)

	sort.OptimizedBubbleSort(arr)
	err := common.CheckSorted(arr)

	if err != nil {
		t.Error(err)
	}
}

func TestSort_CocktailShakerSort(t *testing.T) {
	arr := common.MakeRandomIntSlice(1000, 1000)

	sort.CocktailShakerSort(arr)
	err := common.CheckSorted(arr)

	if err != nil {
		t.Error(err)
	}
}

func TestSort_LomutoPartition(t *testing.T) {
	arr := common.MakeRandomIntSlice(1000, 1000)

	pivot, index := sort.LomutoPartition(arr)

	if pivot != arr[index] {
		t.Error("Partition error: pivot is not at index")
	}

	for i := 0; i < index; i++ {
		if arr[i] > pivot {
			t.Error("Partition error: left side of pivot is not less than pivot")
		}
	}

	for i := index; i < len(arr); i++ {
		if arr[i] < pivot {
			t.Error("Partition error: right side of pivot is not greater than pivot")
		}
	}
}

func TestSort_QuickSort(t *testing.T) {
	arr := common.MakeRandomIntSlice(1000, 1000)

	sort.QuickSort(arr)
	err := common.CheckSorted(arr)

	if err != nil {
		t.Error(err)
	}
}
