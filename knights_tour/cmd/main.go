package main

import (
	"fmt"
	"knightstour"
	"time"
)

// The board dimensions.
const numRows = 6
const numCols = numRows

// Whether we want an open or closed tour.
const requireClosedTour = true

var numCalls int64

func main() {
	board := knightstour.MakeBoard(numRows, numCols)

	startTime := time.Now()
	done := knightstour.StartTour(board, 0, 0, requireClosedTour)

	if done {
		fmt.Printf("Knight's tour found in %v.\n", time.Since(startTime))
	} else {
		fmt.Printf("Knight's tour not found.\n")
	}

	knightstour.DumpBoard(board)
}
