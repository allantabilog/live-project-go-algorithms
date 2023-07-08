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

func placeQueens1(board [][]string, numRows, r, c int) bool {
	return false
}