package recursion

import (
	"math/big"
	"testing"
)

func TestFactorial_Zero(t *testing.T) {
	t.Parallel()

	result := Factorial(0)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestFactorial_Overflow(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Overflow did not panic")
		}
	}()

	Factorial(21)
}

func TestFactorial_Positive(t *testing.T) {
	t.Parallel()

	result := Factorial(5)
	if result != 120 {
		t.Errorf("Expected 120, got %d", result)
	}
}

func TestFactorialBig_Zero(t *testing.T) {
	t.Parallel()

	result := FactorialBig(0)

	want := big.NewInt(1)

	if result.Cmp(want) != 0 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestFactorialBig_One(t *testing.T) {
	t.Parallel()

	result := FactorialBig(1)

	want := big.NewInt(1)

	if result.Cmp(want) != 0 {
		t.Errorf("Expected 1, got %d", result)
	}
}

func TestFactorialBig_Five(t *testing.T) {
	t.Parallel()

	result := FactorialBig(5)

	want := big.NewInt(120)

	if result.Cmp(want) != 0 {
		t.Errorf("Expected 120, got %d", result)
	}
}

func TestFactorialBig_TwentyOne(t *testing.T) {
	t.Parallel()

	result := FactorialBig(21)

	want := new(big.Int)
	want.SetString("51090942171709440000", 10)

	if result.Cmp(want) != 0 {
		t.Errorf("Expected 51090942171709440000, got %d", result)
	}
}
