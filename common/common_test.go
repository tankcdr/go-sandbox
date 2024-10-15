package common_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/tankcdr/common"
)

func TestCommon_MakeRandomIntSlice(t *testing.T) {
	t.Parallel()
	intSlice := common.MakeRandomIntSlice(1000, 1000)

	if len(intSlice) != 1000 {
		t.Errorf("Expected length of 10, but got %d", len(intSlice))
	}
}

func TestCommon_PrintSliceWithSmallerNumItems(t *testing.T) {
	t.Parallel()

	sliceLen := 10
	testSliceLen := 5

	intSlice := common.MakeRandomIntSlice(sliceLen, 1000)

	want, _ := fmt.Fprintln(io.Discard, intSlice[:testSliceLen])

	got, err := common.PrintSlice(intSlice, testSliceLen)

	if err != nil {
		t.Errorf("Error printing slice: %v", err)
	}

	if got != want {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}

func TestCommon_PrintSliceWithLargerNumItems(t *testing.T) {
	t.Parallel()

	sliceLen := 10
	testSliceLen := 20

	intSlice := common.MakeRandomIntSlice(sliceLen, 100)
	want, _ := fmt.Fprintln(io.Discard, intSlice)

	got, err := common.PrintSlice(intSlice, testSliceLen)

	if err != nil {
		t.Errorf("Error printing slice: %v", err)
	}

	if got != want {
		t.Errorf("Expected %d, but got %d", want, got)
	}
}

func TestCommon_CheckSorted(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 3, 4, 5}
	err := common.CheckSorted(slice)

	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}
}

func TestCommon_CheckUnsorted(t *testing.T) {
	t.Parallel()

	slice := []int{1, 2, 3, 5, 4}
	err := common.CheckSorted(slice)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
