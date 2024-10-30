package main

import (
	"fmt"
	"knightstour"
)

// The board dimensions.
const numRows = 8
const numCols = numRows

// Whether we want an open or closed tour.
const requireClosedTour = true

var numCalls int64

func main() {
	board := knightstour.MakeBoard(numRows, numCols)

	closed := knightstour.StartTour(board, 0, 0, requireClosedTour)

	fmt.Printf("Knight's tour found. Meet requirements: %v\n", closed)
	knightstour.DumpBoard(board)
}
