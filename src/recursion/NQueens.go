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
func seriesIsLegal(board [][]string, dim, startRow, startCol, incrRow, incrCol int) bool {
	// maybe validate that incrRow and incrCol are g.t. 0
	// otherwise the loop below will never terminate
	var queenCount = 0

	var row, col = startRow, startCol
	for {
		if row > dim - 1 || col > dim - 1 {
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
	for row := 0; row < dim; row++ {
		for col := 0; col < dim; col++ {
			if !seriesIsLegal(board, dim, row, col, 1, 1) {
				return false
			}
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

	return boardIsLegal && (queenCount == 8)
}



func placeQueens1(board [][]string, dim, row, col int) bool {

	fmt.Println("Board state:")
	Trace(board)


	if row > dim - 1 {
		return boardIsASolution(board, dim)
	}
	var nextRow, nextCol int = row, col + 1

	if nextCol > dim - 1 {
		nextRow = nextRow + 1
	}

	board[row][col] = queen

	if placeQueens1(board, dim, nextRow, nextCol) {
		return true
	}

	board[row][col] = empty

	return false
}