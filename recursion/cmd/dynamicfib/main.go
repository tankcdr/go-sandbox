package main

import (
	"fmt"
	"strconv"

	"github.com/tankcdr/recursion"
)

func main() {
	// Fill-on-the-fly.
	fibonacciValues := make([]int64, 2)
	fibonacciValues[0] = 0
	fibonacciValues[1] = 1

	for {
		// Get n as a string.
		var nString string
		fmt.Printf("N: ")
		fmt.Scanln(&nString)

		// If the n string is blank, break out of the loop.
		if len(nString) == 0 {
			break
		}

		// Convert to int and calculate the Fibonacci number.
		n, _ := strconv.ParseInt(nString, 10, 64)

		// Uncomment one of the following.
		fmt.Printf("fibonacciOnTheFly(%d) = %d\n", n, recursion.DynamicFibonacci(n))
	}

	// Print out all memoized values just so we can see them.
	for i := 0; i < len(fibonacciValues); i++ {
		fmt.Printf("%d: %d\n", i, fibonacciValues[i])
	}
}
