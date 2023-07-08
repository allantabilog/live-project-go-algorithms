package recursion

import (
	"fmt"
)

// the board dimensions
const numRows = 8
const numCols = numRows

// closed or open tour required?
const requireClosedTour = false

// used to mark an unvisited square
const unvisited = -1

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
func initialiseOffsets() {
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
func MakeBoard(numRows int, numCols int) [][]int {
	var board [][]int = make([][]int, numRows)

	for i := range(board) {
		board[i] = make([]int, numCols)
	}
	
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			board[i][j] = unvisited
		}
	}

	return board
}

func Trace(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j :=0 ; j < len(board); j++ {
			fmt.Printf(" %2d", board[i][j])
		}
		fmt.Println()
	}
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
// @todo: increment numCalls
if numVisited == numRows * numCols {
	if !requireClosedTour {
		return true
	}
	// check to see if this is a closed tour
	// try every move from here and if any of these moves 
	// land back on the starting position - we have a closed tour (return true)
	// otherwise - not a closed tour, return false
	
}
// not every square has been visited
// continue searching for a tour from here

// @todo remove this when done
return false

}

// Check that a move to (row, col) on the board is allowed
func ValidateMove(toRow int, toCol int, board [][]int) bool {
	if toRow < 0 || toRow > 7 || toCol < 0 || toCol > 7 {
		return false
	}
	if board[toRow][toCol] != unvisited {
		return false
	}
	return true
}
