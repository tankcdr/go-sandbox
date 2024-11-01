package nqueens

import "fmt"

const EMPTY = "."
const QUEEN = "Q"

// Create a board of size m x m
func MakeBoard(m int) [][]string {
	board := make([][]string, m)
	for i := range board {
		board[i] = make([]string, m)
		for j := range board[i] {
			board[i][j] = EMPTY
		}
	}
	return board
}

func DumpBoard(board [][]string) {
	for i := range board {
		for j := range board[i] {
			fmt.Printf("%s ", board[i][j])
		}
		println()
	}
}

func NQueens(board [][]string, r, c int) bool {
	max := len(board)

	if r >= max {
		return boardIsASolution(board)
	}

	// Find the next square.
	nextR := r
	nextC := c + 1
	if nextC >= max {
		nextR += 1
		nextC = 0
	}

	if NQueens(board, nextR, nextC) {
		return true
	}

	board[r][c] = QUEEN
	if boardIsLegal(board) {
		if NQueens(board, nextR, nextC) {
			return true
		}
	}
	board[r][c] = EMPTY

	return false
}

// Checks if the current state a solution
func boardIsASolution(board [][]string) bool {
	if boardIsLegal(board) {
		return countQueens(board) == len(board)
	}
	return false
}

// Count the number of queens on the board
func countQueens(board [][]string) int {
	count := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == QUEEN {
				count++
			}
		}
	}
	return count
}

// Check if the board is legal (no queens attacking each other)
func boardIsLegal(board [][]string) bool {
	//assumes m x m board
	max := len(board)

	//check rows
	for i := 0; i < max; i++ {
		if !seriesIsLegal(board, i, 0, 0, 1) {
			return false
		}
	}

	//check columns
	for i := 0; i < max; i++ {
		if !seriesIsLegal(board, 0, i, 1, 0) {
			return false
		}
	}

	//check diagonals
	for i := 0; i < max; i++ {
		if !seriesIsLegal(board, i, 0, 1, 1) {
			return false
		}
		if !seriesIsLegal(board, 0, i, 1, 1) {
			return false
		}
		if !seriesIsLegal(board, i, max-1, 1, -1) {
			return false
		}
		if !seriesIsLegal(board, 0, i, 1, -1) {
			return false
		}
	}

	return true
}

func seriesIsLegal(board [][]string, r0, c0, dr, dc int) bool {
	//assumes m x m board
	max := len(board)
	count := 0

	for r0 >= 0 && r0 < max && c0 >= 0 && c0 < max {
		if board[r0][c0] == QUEEN {
			count++
			if count > 1 {
				return false
			}
		}
		r0 += dr
		c0 += dc
	}

	return true
}
