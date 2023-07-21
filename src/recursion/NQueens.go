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


func seriesIsLegal(board [][]string, dim, startRow, startCol, incrRow, incCol int) bool {
    hasQueen := false

    r := startRow
    c := startCol
    for {
        if board[r][c] == queen {
            // If we already have a queen on this row,
            // then this board is not legal.
            if hasQueen { return false }

            // Remember that we have a queen on this row.
            hasQueen = true
        }

        // Move to the next square in the series.
        r += incrRow
        c += incCol

        // If we fall off the board, then the series is legal.
        if  r >= dim ||
            c >= dim ||
            r < 0 ||
            c < 0 {
                return true
        }
    }
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
		if !seriesIsLegal(board, dim, row, 0, 1, 1) {
			return false
		}
	}
	for col := 0; col < dim; col++ {
		if !seriesIsLegal(board, dim, 0, col, 1, 1) {
			return false
		}
	}
	for row := 0; row < dim; row++ {
		if !seriesIsLegal(board, dim, row, dim - 1, 1, -1) {
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
	Trace(board)

	if row >= dim {
		return boardIsASolution(board, dim)
	}

	board[row][col] = queen
	 if boardIsLegal(board, dim) {
		//good. make a recursive call to try the next square
		nextRow, nextCol := row, col + 1
		if nextCol >= dim {
			nextRow = row + 1
			nextCol = 0
		}
		return placeQueens1(board, dim, nextRow, nextCol)
	 } else {
		// backtrack
		board[row][col] = empty
		nextRow, nextCol := row, col + 1
		if nextCol >= dim {
			nextRow = row + 1
			nextCol = 0
		}
		return placeQueens1(board, dim, nextRow, nextCol)
	 }
}


// func placeQueens1(board [][]string, dim, row, col int) bool {

// 	fmt.Println("Board state:")
// 	Trace(board)


// 	if row >= dim {
// 		// We have made a complete assignment of queens
// 		// check if we have a solution
// 		return boardIsASolution(board, dim)
// 	}
// 	var nextRow, nextCol int = row, col + 1

// 	if nextCol >= dim {
// 		// we have fallen off the board
// 		// move down to the beginning of the next row
// 		nextRow = nextRow + 1
// 		nextCol = 0
// 	}

// 	// // what happens if we do not place a square in row, col?
// 	// // examine the next square 
// 	// if placeQueens1(board, dim, nextRow, nextCol) {
// 	// 	return true
// 	// } 

// 	board[row][col] = queen
// 	if placeQueens1(board, dim, nextRow, nextCol) {
// 		return true
// 	} else {
// 		// we backtrack
// 		board[row][col] = empty
// 		return false
// 	}
// }