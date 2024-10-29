package recursion

func Fibonacci(n int64) int64 {
	if n < 0 {
		panic("n must be a non-negative integer")
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
