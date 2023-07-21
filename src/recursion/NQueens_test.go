package recursion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrive(t *testing.T) {
	NQueensDriver()
}

func TestSeriesIsLegal_Rows(t *testing.T) {
	// generate some test boards
	var boardGood = MakeBoard[string](10, 10, empty)
	boardGood[0][0] = queen
	boardGood[1][0] = queen
	Trace(boardGood)

	var boardBad = MakeBoard[string](10, 10, empty)
	boardBad[0][0] = queen
	boardBad[0][3] = queen
	Trace(boardBad)

	// boardGood passes row series test (no adjacent queens in same row)
	assert.Equal(t, true, seriesIsLegal(boardGood, 10, 0, 0, 0, 1), "Board has < 2 queens in a row")

	// boardBad fails row series 
	assert.Equal(t, false, seriesIsLegal(boardBad, 10, 0, 0, 0, 1), "Board has >= 2 queens in a row")
}

func TestSeriesIsLegal_Columns(t *testing.T) {
	// generate some test boards
	var boardGood = MakeBoard[string](10, 10, empty)
	boardGood[0][0] = queen
	boardGood[0][1] = queen
	Trace(boardGood)

	var boardBad = MakeBoard[string](10, 10, empty)
	boardBad[0][0] = queen
	boardBad[1][0] = queen
	Trace(boardBad)

	// boardGood passes column series
	assert.Equal(t, true, seriesIsLegal(boardGood, 10, 0, 0, 1, 0), "Board has < 2 queens in a column")

	// boardBad fails column series 
	assert.Equal(t, false, seriesIsLegal(boardBad, 10, 0, 0, 1, 0), "Board has >= 2 queens in a column")
}

func TestSeriesIsLegal_Diagonal(t *testing.T) {
	// generate some test boards
	var boardGood = MakeBoard[string](10, 10, empty)
	boardGood[0][0] = queen
	boardGood[0][1] = queen
	Trace(boardGood)

	var boardBad = MakeBoard[string](10, 10, empty)
	boardBad[0][0] = queen
	boardBad[1][1] = queen
	Trace(boardBad)

	// boardGood passes column series
	assert.Equal(t, true, seriesIsLegal(boardGood, 10, 0, 0, 1, 1), "Board has < 2 queens in a diagonal")

	// boardBad fails column series 
	assert.Equal(t, false, seriesIsLegal(boardBad, 10, 0, 0, 1, 1), "Board has >= 2 queens in a diagonal")
}

func TestBoardIsLegal(t *testing.T) {
	// generate some test boards
	var emptyBoard = MakeBoard[string](4, 4, empty)

	var goodBoard = MakeBoard[string](4, 4, empty)
	goodBoard[0][1] = queen
	goodBoard[1][3] = queen
	goodBoard[2][0] = queen
	goodBoard[3][2] = queen

	fmt.Println("goodBoard")
	Trace(goodBoard)

	var badBoard = MakeBoard[string](4, 4, empty)
	badBoard[0][0] = queen
	badBoard[0][1] = queen
	badBoard[1][0] = queen
	badBoard[1][1] = queen

	fmt.Println("badBoard")
	Trace(badBoard)


	var testCases = []struct {
		board [][]string
		expected bool
		description string
	}{
		{board: emptyBoard, expected: true, description: "empty board is legal" },
		{board: goodBoard, expected: true, description: "good board is legal" },
		{board: badBoard, expected: false, description: "bad board is not legal" },
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, boardIsLegal(tc.board, 4), tc.description)
	}

}

func TestBoardIsLegal_bad(t *testing.T) {
	var badBoard = MakeBoard[string](3, 3, empty)
	badBoard[0][0] = queen
	badBoard[1][2] = queen
	badBoard[2][1] = queen

	Trace(badBoard)

	assert.Equal(t, false, boardIsLegal(badBoard, 3))

}


func TestQueenCount(t *testing.T) {
	var emptyBoard = MakeBoard[string](10, 10, empty)

	var nonEmptyBoard = MakeBoard[string](10, 10, empty)
	nonEmptyBoard[0][0] = queen
	nonEmptyBoard[1][2] = queen

	assert.Equal(t, 0, queenCount(emptyBoard, 10))
	assert.Equal(t, 2, queenCount(nonEmptyBoard, 10))
}

func TestBoardIsASolution(t *testing.T){
	var emptyBoard = MakeBoard[string](8, 8, empty)

	var smallBoard = MakeBoard[string](8, 8, empty)
	smallBoard[0][0] = queen

	var goodBoard = MakeBoard[string](8, 8, empty)
	goodBoard[0][0] = queen
	goodBoard[1][2] = queen
	goodBoard[2][4] = queen
	goodBoard[3][6] = queen
	goodBoard[5][3] = queen
	goodBoard[6][1] = queen

	Trace(goodBoard)
	assert.True(t, boardIsLegal(goodBoard, 8), "goodBoard is legal")
	assert.False(t, boardIsASolution(goodBoard, 8), "goodBoard is legal")

	assert.False(t, boardIsASolution(emptyBoard, 8))
	assert.False(t, boardIsASolution(smallBoard, 8))
}

func TestPlaceQueens1_size3(t *testing.T){
	var board = MakeBoard[string](3, 3, empty)
	var solution = placeQueens1(board, 3, 0, 0)
	fmt.Println("Final board state: ")
	Trace(board)
	fmt.Println("Solved? ", solution)
}

func TestPlaceQueens1_size4(t *testing.T){
	var board = MakeBoard[string](4, 4, empty)
	var solution = placeQueens1(board, 4, 0, 0)
	fmt.Println("Final board state: ")
	Trace(board)
	fmt.Println("Solved? ", solution)
}

func TestPlaceQueens1_size4_attempt2(t *testing.T){
	var board = MakeBoard[string](4, 4, empty)
	var solution = placeQueens1(board, 4, 0, 2)
	fmt.Println("Final board state: ")
	Trace(board)
	fmt.Println("Solved? ", solution)
}

func TestPlaceQueens1_size5(t *testing.T){
	var board = MakeBoard[string](5, 5, empty)
	var solution = placeQueens1(board, 5, 0, 0)
	fmt.Println("Final board state: ")
	Trace(board)
	fmt.Println("Solved? ", solution)
}
func TestPlaceQueens1_size6(t *testing.T){
	var board = MakeBoard[string](6, 6, empty)
	var solution = placeQueens1(board, 6, 0, 5)
	fmt.Println("Final board state: ")
	Trace(board)
	fmt.Println("Solved? ", solution)
}
func TestPlaceQueens1_size7(t *testing.T){
	var board = MakeBoard[string](7, 7, empty)
	var solution = placeQueens1(board, 7, 0, 0)
	fmt.Println("Final board state: ")
	Trace(board)
	fmt.Println("Solved? ", solution)
}