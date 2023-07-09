package recursion

import (
	"fmt"
	"time"
)

const (
	empty = "."
	queen = "Q"
	qnNumRows = 10
	qnNumCols = 10
)


func NQueensDriver() {
	board := MakeBoard(qnNumRows, qnNumCols, empty)	
	Trace(board)

	start := time.Now()
	success := placeQueens1(board, 10, 0, 0)

	elapsed := time.Since(start)

	if success {
		fmt.Println("Success")
		Trace(board)
	} else {
		fmt.Println("No solution.")
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

}

// Returns true if the given series of squares contains at most one queen
func seriesIsLegal(board [][]string, numRows, startRow, startCol, incrRow, incrCol int) bool {
	// maybe validate that incrRow and incrCol are g.t. 0
	// otherwise the loop below will never terminate
	var queenCount = 0

	var row, col = startRow, startCol
	for {
		if row > qnNumRows - 1 || col > qnNumCols - 1 {
			break
		}
		if board[row][col] == queen {
			queenCount ++
		}
		row, col = row + incrRow, col + incrCol
	}

	return queenCount < 2
}

// returns true iff
// every row, column, and diagonal on the board is legal
func boardIsLegal(board [][]string, dim int) bool {
	// scan all rows
	for row := 0; row < dim; row++ {
		if !seriesIsLegal(board, 10, row, 0, 0, 1) {
			return false
		}
	}
	
	// scan all columns
	for col := 0; col < dim; col++ {
		if !seriesIsLegal(board, 10, 0, col, 1, 0) {
			return false
		}
	}
	// scan all diagonals
	// @todo. find an algorithm for scanning all diagonals systematically
	if !seriesIsLegal(board, 10, 0, 0, 1, 1) {
		return false
	}
	return true
}
func placeQueens1(board [][]string, numRows, r, c int) bool {
	return false
}