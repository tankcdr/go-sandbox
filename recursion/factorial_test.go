package recursion

import "testing"

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
