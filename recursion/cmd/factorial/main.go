package main

import (
	"fmt"

	"math"

	"github.com/tankcdr/recursion"
)

func main() {

	fmt.Printf("max uint64: %v\n", uint64(math.MaxUint64))
	var n uint64
	for n = 0; n <= 21; n++ {
		fmt.Printf("%3d! = %20d\n", n, recursion.Factorial(n))
	}
	fmt.Println()
}
