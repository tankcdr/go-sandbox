package recursion

import (
	"math"
	"math/big"
)

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

func FactorialBig(n uint64) *big.Int {
	if n == 0 || n == 1 {
		return big.NewInt(1)
	}

	result := FactorialBig(n - 1)

	return result.Mul(result, big.NewInt(int64(n)))
}
