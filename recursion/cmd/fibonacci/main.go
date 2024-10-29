// Fibonacci numbers
package main

import (
	"fmt"
	"strconv"

	"github.com/tankcdr/recursion"
)

func main() {
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
		fmt.Printf("fibonacci(%d) = %d\n", n, recursion.Fibonacci(n))
	}
}
