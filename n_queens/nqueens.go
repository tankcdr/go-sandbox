package nqueens

import (
	"fmt"
)

const EMPTY = "."
const QUEEN_VALUE = "Q"

const FREE = 0
const ATTACKED = 1
const QUEEN = 1000000

// Create a board of size m x m
func MakeBoard(m int) [][]int {
	board := make([][]int, m)
	for i := range board {
		board[i] = make([]int, m)
		for j := range board[i] {
			board[i][j] = FREE
		}
	}
	return board
}

func DumpBoard(board [][]int) {
	for i := range board {
		for j := range board[i] {
			switch board[i][j] {
			case QUEEN:
				fmt.Printf("%2s ", QUEEN_VALUE)
				break
			default:
				fmt.Printf("%2s ", EMPTY)

			}
		}
		println()
	}
}

// Solve the N-Queens problem by placing queens by row
func NQueens(board [][]int, r, c, placed int) bool {
	max := len(board)

	if r >= max {
		return boardIsASolution(board, placed)
	}

	// Find the next square.
	nextR := r
	nextC := c + 1
	if nextC >= max {
		nextR += 1
		nextC = 0
	}

	if NQueens(board, nextR, nextC, placed) {
		return true
	}

	if board[r][c] == FREE {

		placeQueen(board, r, c, max)

		//if boardIsLegal(board) {
		if NQueens(board, nextR, nextC, placed+1) {
			return true
		}
		//}

		removeQueen(board, r, c, max)
	}

	return false
}

// Solve the N-Queens problem by placing queens by column
// This is a recursive function that places a queen in the current column
// and then recursively tries to place a queen in the next column
func NQueensByColumn(board [][]int, numRows, c int) bool {

	isLegal := boardIsLegal(board)

	if c >= numRows {
		return isLegal
	} else if !isLegal {
		return isLegal
	}

	for r := 0; r < numRows; r++ {
		if board[r][c] == FREE {
			placeQueen(board, r, c, numRows)
			if NQueensByColumn(board, numRows, c+1) {
				return true
			}
			removeQueen(board, r, c, numRows)
		}
	}

	return false
}

// Place a queen on the board and update the board to reflect the new queen placement
func placeQueen(board [][]int, r, c, max int) {
	updateBoard(board, r, c, max, ATTACKED)
	board[r][c] = QUEEN
}

func removeQueen(board [][]int, r, c, max int) {
	updateBoard(board, r, c, max, -ATTACKED)
	board[r][c] = FREE
}

// Update the board to reflect the new queen placement and her impact on the board
// or remove the queen and her impact from the board
func updateBoard(board [][]int, r, c, max, value int) {
	//update row
	adjustAttacks(board, r, 0, 0, 1, max, value)

	//update columns
	adjustAttacks(board, 0, c, 1, 0, max, value)

	//update diagonals

	adjustAttacks(board, r, c, 1, 1, max, value)
	adjustAttacks(board, r, c, -1, 1, max, value)
	adjustAttacks(board, r, c, 1, -1, max, value)
	adjustAttacks(board, r, c, -1, -1, max, value)

}

// Adjust the attacks on the board
// Updates the section of the board (row, col, diagonal) that the queen attacks
// or clears the attack if the queen is removed
func adjustAttacks(board [][]int, r0, c0, dr, dc, max, value int) {
	for r0 >= 0 && r0 < max && c0 >= 0 && c0 < max {
		if board[r0][c0] != QUEEN {
			board[r0][c0] += value
		}
		r0 += dr
		c0 += dc
	}

}

// Checks if the current state a solution
func boardIsASolution(board [][]int, placed int) bool {
	//optimizing by remove countQueens
	if boardIsLegal(board) {
		return placed >= len(board)
		//return countQueens(board) == len(board)
	}
	return false

}

// Count the number of queens on the board
/*
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
}*/

// Check if the board is legal (no queens attacking each other)
func boardIsLegal(board [][]int) bool {
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

// Check if a series of squares is legal (no queens attacking each other)
// series = row, column, or diagonal depending upon the parameters
// dr, dc = direction to move in the series
func seriesIsLegal(board [][]int, r0, c0, dr, dc int) bool {
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
