package knightstour

import (
	"fmt"
	"math/rand/v2"
)

// Value to represent a square that we have not visited.
const UNVISITED = -1

// Define offsets for the knight's movement.
type Offset struct {
	Row, Col int
}

var moveOffsets = []Offset{
	{-2, 1},
	{-2, -1},
	{2, 1},
	{2, -1},
	{1, 2},
	{1, -2},
	{-1, 2},
	{-1, -2},
}

// Create a board of size m x n
func MakeBoard(m, n int) [][]int {
	board := make([][]int, m)
	for i := range board {
		board[i] = make([]int, n)
		for j := range board[i] {
			board[i][j] = UNVISITED
		}
	}
	return board
}

func DumpBoard(board [][]int) {
	for i := range board {
		for j := range board[i] {
			fmt.Printf("%02d ", board[i][j])
		}
		println()
	}
}

func StartTour(board [][]int, startingRow int, startingCol int, requiredClosedTour bool) bool {
	numRows := len(board)
	numCols := len(board[0])

	//starting point
	board[startingRow][startingCol] = 0

	found := FindTour(board, numRows, numCols, startingRow, startingCol, 0, requiredClosedTour)

	return found
}

func FindTour(board [][]int, numRows int, numCols int, currentRow int, currentCol int, numVisited int, requiredClosed bool) bool {
	numVisited++

	// Check if we have visited all the squares
	if numVisited == numRows*numCols {

		if !requiredClosed {
			return true
		}

		// Check if we can return to the starting point
		for _, offset := range moveOffsets {
			newRow := currentRow + offset.Row
			newCol := currentCol + offset.Col

			if newRow == 0 && newCol == 0 {
				return true
			}
		}
		return false
	}

	// Randomize the order of the moves
	rand.Shuffle(len(moveOffsets), func(i, j int) { moveOffsets[i], moveOffsets[j] = moveOffsets[j], moveOffsets[i] })

	//walk through all possible moves
	for _, offset := range moveOffsets {
		newRow := currentRow + offset.Row
		newCol := currentCol + offset.Col

		if isValidMove(board, numRows, numCols, newRow, newCol) {
			board[newRow][newCol] = numVisited

			if FindTour(board, numRows, numCols, newRow, newCol, numVisited, requiredClosed) {
				return true
			}

			board[newRow][newCol] = UNVISITED
		}
	}

	return false
}

func isValidMove(board [][]int, numRows int, numCols int, newRow int, newCol int) bool {
	return newRow >= 0 && newRow < numRows && newCol >= 0 && newCol < numCols && board[newRow][newCol] == UNVISITED
}
