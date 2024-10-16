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
