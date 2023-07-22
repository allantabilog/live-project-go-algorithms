package main

import (
	"fmt"
	"time"
)

const (
	empty = "."
	queen = "Q"
)

func makeBoard(dim int) [][]string {
	board := make([][]string, dim)
	for r := range board {
		board[r] = make([]string, dim)
		for c := 0; c < dim; c++ {
			board[r][c] = empty
		}
	}
	return board
}

func trace(board [][]string) {
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			fmt.Printf("%2s", board[r][c])
		}
		fmt.Println()
	}
}

// Rod's code: a series is a single row, a single column, or a single diagonal
// idea: as soon as you find two queens in a series, you're done - return false
func seriesIsLegal(board [][]string, dim, startRow, startCol, incrRow, incCol int) bool {
	hasQueen := false // this flag gets set as soon as we find a queen on a square

	r := startRow
	c := startCol
	for {
		if board[r][c] == queen {
			// If we already have a queen on this row,
			// then this board is not legal.
			if hasQueen {
				return false
			}

			// Remember that we have a queen on this row.
			hasQueen = true
		}

		// Move to the next square in the series.
		r += incrRow
		c += incCol

		// If we fall off the board, then the series is legal.
		if r >= dim ||
			c >= dim ||
			r < 0 ||
			c < 0 {
			return true
		}
	}
}

func boardIsLegal(board [][]string, dim int) bool {
	// scan all rows
	for row := 0; row < dim; row++ {
		if !seriesIsLegal(board, dim, row, 0, 0, 1) {
			return false
		}
	}

	// scan all columns
	for col := 0; col < dim; col++ {
		if !seriesIsLegal(board, dim, 0, col, 1, 0) {
			return false
		}
	}
	// scan all diagonals
	// diagonals down to the right
	for row := 0; row < dim; row++ {
		if !seriesIsLegal(board, dim, row, 0, 1, 1) {
			return false
		}
	}
	for col := 0; col < dim; col++ {
		if !seriesIsLegal(board, dim, 0, col, 1, 1) {
			return false
		}
	}
	// diagonals down to the left
	for row := 0; row < dim; row++ {
		if !seriesIsLegal(board, dim, row, dim-1, 1, -1) {
			return false
		}
	}
	for col := 0; col < dim; col++ {
		if !seriesIsLegal(board, dim, 0, col, 1, -1) {
			return false
		}
	}

	return true
}

func queenCount(board [][]string, dim int) int {
	var queenCount = 0

	for row := 0; row < dim; row++ {
		for col := 0; col < dim; col++ {
			if board[row][col] == queen {
				queenCount++
			}
		}
	}
	return queenCount
}
func boardIsASolution(board [][]string, dim int) bool {

	var boardIsLegal bool = boardIsLegal(board, dim)
	var queenCount int = queenCount(board, dim)

	return boardIsLegal && (queenCount == dim)
}

func placeQueens1(board [][]string, dim, row, col int) bool {

	if row >= dim {
		return boardIsASolution(board, dim)
	}

	nextRow, nextCol := row, col+1

	if nextCol >= dim {
		nextRow = nextRow + 1
		nextCol = 0
	}

	// don't place a queen here and check the next square
	if placeQueens1(board, dim, nextRow, nextCol) {
		return true
	}

	// place a queen here and check the next square
	board[row][col] = queen
	if placeQueens1(board, dim, nextRow, nextCol) {
		return true
	}

	// backtrack to undo last move
	board[row][col] = empty

	// return false and backtrack further
	return false
}

func placeQueens2(board [][]string, dim, queensPlaced, row, col int) bool {

	if queensPlaced == dim {
		return boardIsASolution(board, dim)
	}

	if row >= dim {
		return false
	}

	nextRow, nextCol := row, col+1

	if nextCol >= dim {
		nextRow = nextRow + 1
		nextCol = 0
	}

	// don't place a queen here and check the next square
	if placeQueens2(board, dim, queensPlaced, nextRow, nextCol) {
		return true
	}

	// place a queen here and check the next square
	board[row][col] = queen
	queensPlaced += 1
	if placeQueens2(board, dim, queensPlaced, nextRow, nextCol) {
		return true
	}

	// backtrack to undo last move
	board[row][col] = empty
	queensPlaced -= 1

	// return false and backtrack further
	return false
}

func placeQueens4(board [][]string, dim, col int) bool {

	if col == dim {
		return boardIsLegal(board, dim)
	}

	if !boardIsLegal(board, dim) {
		return false
	}

	for row := 0; row < dim; row++ {
		nextCol := col + 1
		board[row][col] = queen
		if placeQueens4(board, dim, nextCol) {
			return true
		}
		board[row][col] = empty
	}
	return false
}

func main() {
	// try(2)
	// try(3)
	// try(4)
	// try(5)
	// try(6)
	// try(7)
	// try(8)
	// try(9)
	// try(10)
	try(20)
}

func try(dim int) {
	board := makeBoard(dim)
	start := time.Now()
	success := placeQueens4(board, dim, 0)
	elapsed := time.Since(start)

	if success {
		fmt.Printf("Found a solution for n=%d\n", dim)
		trace(board)
	} else {
		fmt.Printf("Did not find a solution for n=%d\n", dim)
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}
