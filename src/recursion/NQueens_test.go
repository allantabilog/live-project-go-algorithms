package recursion

import (
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
	var emptyBoard = MakeBoard[string](10, 10, empty)

	var goodBoard = MakeBoard[string](10, 10, empty)
	goodBoard[0][0] = queen
	goodBoard[1][3] = queen
	goodBoard[2][4] = queen
	goodBoard[3][7] = queen
	goodBoard[9][1] = queen
	goodBoard[5][6] = queen
	goodBoard[4][8] = queen
	goodBoard[6][9] = queen

	var badBoard = MakeBoard[string](10, 10, empty)
	badBoard[0][0] = queen
	badBoard[1][3] = queen
	badBoard[2][4] = queen
	badBoard[3][7] = queen
	badBoard[9][1] = queen
	badBoard[5][6] = queen
	badBoard[4][8] = queen
	badBoard[6][8] = queen

	Trace(goodBoard)


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
		assert.Equal(t, tc.expected, boardIsLegal(tc.board, 10), tc.description)
	}

}