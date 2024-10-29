package recursion_test

import (
	"testing"

	"github.com/tankcdr/recursion"
)

func TestFibonacci_Negative(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	recursion.Fibonacci(-1)
}

func TestFibonacci_Zero(t *testing.T) {
	t.Parallel()

	got := recursion.Fibonacci(0)
	want := int64(0)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFibonacci_One(t *testing.T) {
	t.Parallel()

	got := recursion.Fibonacci(1)
	want := int64(1)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFibonacci(t *testing.T) {
	t.Parallel()

	got := recursion.Fibonacci(10)
	want := int64(55)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
