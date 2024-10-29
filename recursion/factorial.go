package recursion

import "math"

// Factorial calculates the factorial of a given number
func Factorial(n uint64) uint64 {
	if n == 0 || n == 1 {
		return 1
	}

	result := Factorial(n - 1)

	if result > math.MaxUint64/n {
		panic("Factorial result is too large to fit in a uint64")
	}

	return result * n
}
