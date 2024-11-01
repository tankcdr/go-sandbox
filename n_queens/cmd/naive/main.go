package main

import (
	"fmt"
	"nqueens"
	"time"
)

func main() {
	const numRows = 13
	board := nqueens.MakeBoard(numRows)

	start := time.Now()
	success := nqueens.NQueens(board, 0, 0, 0)
	elapsed := time.Since(start)

	if success {
		fmt.Println("Success!")
		nqueens.DumpBoard(board)
	} else {
		fmt.Println("No solution")
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}
