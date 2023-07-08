package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnightsTour(t *testing.T) {
	var board = MakeBoard(8, 8)
	InitialiseOffsets()
	
	fmt.Println("Before tour search:")
	Trace(board)
	
	var result = FindTour(board, 8, 8, 0, 0, 0)
	
	if result {
		fmt.Println("Successfully found a tour")
	} else {
		fmt.Println("Failed to find a tour")
	}
	fmt.Println("After tour search:")
	Trace(board)
}

func TestValidateMove(t *testing.T) {
	var board = MakeBoard(8, 8)
	// setup an initial configuration
	var row, col = 4, 4
	board[row][col] = 0
	
	row, col = row - 1, col - 2
	board[row][col] = 1

	row, col = row - 2, col - 1
	board[row][col] = 2

	// hence positions (4, 4), (3, 2), and (1, 1)
	// are visited and should be flagged as invalid
	Trace(board)

	testData := [] struct {
		toRow int 
		toCol int 
		expected bool
	}{
		{0, 0, true},
		{1, 2, true},
		{4, 4, false},
		{3, 2, false},
		{1, 1, false},
		{-1, 1, false},
		{1, 8, false},
		{9, 9, false},
	}
	for _, tc := range testData {
		assert.Equal(t, tc.expected, ValidateMove(tc.toRow, tc.toCol, board))
	}
}