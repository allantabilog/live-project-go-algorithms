package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrive(t *testing.T) {
	NQueensDriver()
}

func TestSeriesIsLegal(t *testing.T) {
	board := MakeBoard[string](10, 10, empty)
	Trace(board)


	// generate some test boards
	var boardGood = MakeBoard[string](10, 10, empty)
	boardGood[0][0] = queen
	boardGood[2][3] = queen

	var boardBad = MakeBoard[string](10, 10, empty)
	boardBad[0][0] = queen
	boardBad[0][3] = queen

	// boardGood passes row series test (no adjacent queens in same row)
	assert.Equal(t, true, seriesIsLegal(boardGood, 10, 0, 0, 1, 0))
	// boardGood passes column series test (no adjacent queens in same col)
	assert.Equal(t, true, seriesIsLegal(boardGood, 10, 0, 0, 0, 1))
	// boardBad passes column series test (no adjacent queens in same col)
	assert.Equal(t, true, seriesIsLegal(boardGood, 10, 0, 0, 0, 1))
	// boardBad fails row series test (no adjacent queens in same row)
	assert.Equal(t, true, seriesIsLegal(boardGood, 10, 0, 0, 1, 0))
	

}