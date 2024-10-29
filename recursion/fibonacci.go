package recursion

// Fibonacci calculates the nth Fibonacci number using recursion
// naive implementation
func Fibonacci(n int64) int64 {
	if n < 0 {
		panic("n must be a non-negative integer")
	}

	if n > 92 {
		panic("n must be less than or equal to 92, otherwise the result will overflow")
	}

	return fibonacci(n)
}

func fibonacci(n int64) int64 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// DynamicFibonacci calculates the nth Fibonacci number using recursion
// with memory to avoid recalculating the same values
func DynamicFibonacci(n int64) int64 {
	if n < 0 {
		panic("n must be a non-negative integer")
	}

	if n > 92 {
		panic("n must be less than or equal to 92, otherwise the result will overflow")
	}

	memory := make(map[int64]int64, 93)
	memory[0] = 0
	memory[1] = 1

	return dynamicFibonacci(n, memory)
}

func dynamicFibonacci(n int64, mem map[int64]int64) int64 {

	if _, exists := mem[n]; !exists {
		mem[n] = dynamicFibonacci(n-1, mem) + dynamicFibonacci(n-2, mem)
	}

	return mem[n]
}
