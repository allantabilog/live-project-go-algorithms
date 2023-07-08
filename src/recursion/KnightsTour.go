package recursion

import (
	"fmt"
)

const (
	// the board dimensions
	numRows = 8
	numCols = 8

	// closed or open tour required?
	requireClosedTour = false
	
	// used to mark an unvisited square
	unvisited = -1
)



// defines "offsets" for a knight's legal move relative 
/// to its current position
// dr - number of row "steps" e.g. +2 means two rows down, -1 one row up
// dc - number of column "steps" e.g. +1 means 1 col right, -2 two cols left
type Offset struct {
	dr, dc int
}

// slice to hold all the possible legal moves that a knight can make
var moveOffsets []Offset

// global var to keep track of number of calls
// not necessary for the implementation
var numCalls int64

// this initialises moveOffsets with offsets
// relative to the current position i.e.
// note...depending on the current position, some of these moves
// are not legal
func InitialiseOffsets() {
	moveOffsets = []Offset{
		{-2, -1},
		{-1, -2},
		{+2, -1},
		{+1, -2},
		{-2, +1},
		{-1, +2},
		{+2, +1},
		{+1, +2},
	}
}

// Function to construct the board with dimensions
// numRows x numCols
// each square initialised to -1 (unvisited)
// board[i][j] == v means that:
// the knight visited square (i,j) in step v
// the starting square gets the value 0
func MakeBoard[T comparable](numRows int, numCols int, initialValue T) [][]T {
	var board [][]T = make([][]T, numRows)

	for i := range(board) {
		board[i] = make([]T, numCols)
	}
	
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			board[i][j] = initialValue
		}
	}

	return board
}

func Trace[T comparable](board [][]T) {
	for i := 0; i < len(board); i++ {
		for j :=0 ; j < len(board); j++ {
			fmt.Printf(" %v", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

// The main function to find a tour if it exists or not
// ?question: which tour?
func FindTour(
	board [][]int, 
	numRows int,
	numCols int,
	curRow int,
	curCol int,
	numVisited int,
) bool {

// store the previous value before updating it
// we will use this to check for closed tours i.e.
// if the last position is the same as the first
var prevVisitedValue = board[curRow][curCol]	

// mark the current position with the visit number
board[curRow][curCol] = numVisited

if numVisited == numRows * numCols {
	if !requireClosedTour {
		return true
	}
	// if the last position is the same as the first then we have a closed tour
	// otherwise, we don't have a closed tour
	if prevVisitedValue == board[curRow][curCol] {
		return true
	}
	return false
}
// not every square has been visited
// continue searching for a tour from here
var nextRow, nextCol = FindNextMove(curRow, curCol, board)

if nextRow != -1 && nextCol != -1 {
	FindTour(board, numRows, numCols, nextRow, nextCol, numVisited + 1)
} else {
	fmt.Println("No more possible moves")
	return false
}


// @todo remove this when done
return false

}

// Work out the next position from current position
// can use different strategies
// this is the simplest one: just try every offset in the list
// and stick with the first one that is valid
func FindNextMove(currRow int, currCol int, board [][]int) (int, int) {
	var nextRow, nextCol = -1, -1
	for _, move := range moveOffsets {
		var testRow, testCol = currRow + move.dr, currCol + move.dc
		if ValidateMove(testRow, testCol, board) {
			nextRow, nextCol = testRow, testCol
			break
		} 
	}
	return nextRow, nextCol
}

// Check that a move to (row, col) on the board is allowed
func ValidateMove(toRow int, toCol int, board [][]int) bool {
	if toRow < 0 || toRow > numRows - 1 || toCol < 0 || toCol > numCols - 1 {
		return false
	}
	if board[toRow][toCol] != unvisited {
		return false
	}
	return true
}
