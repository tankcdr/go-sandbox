package main

import (
	"fmt"
	"nqueens"
	"time"
)

func main() {
	const numRows = 27
	board := nqueens.MakeBoard(numRows)

	start := time.Now()
	success := nqueens.NQueensByColumn(board, numRows, 0)
	elapsed := time.Since(start)

	if success {
		fmt.Println("Success!")
		nqueens.DumpBoard(board)
	} else {
		fmt.Println("No solution")
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}
