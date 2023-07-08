package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrive(t *testing.T) {
	NQueensDriver()
}

func TestSeriesIsLegal(t *testing.T) {

	// generate some test boards
	var boardGood = MakeBoard[string](10, 10, empty)
	boardGood[0][0] = queen
	boardGood[2][3] = queen
	Trace(boardGood)

	var boardBad = MakeBoard[string](10, 10, empty)
	boardBad[0][0] = queen
	boardBad[0][3] = queen
	Trace(boardBad)

	var boardBad2 = MakeBoard[string](10, 10, empty)
	boardBad2[0][0] = queen
	boardBad2[1][1] = queen
	Trace(boardBad2)


	// boardGood passes row series test (no adjacent queens in same row)
	assert.Equal(t, true, seriesIsLegal(boardGood, 10, 0, 0, 1, 0))

	// boardGood passes column series test (no adjacent queens in same col)
	assert.Equal(t, true, seriesIsLegal(boardGood, 10, 0, 0, 0, 1))
	
	// boardBad passes row series
	assert.Equal(t, true, seriesIsLegal(boardBad, 10, 0, 0, 1, 0))

	// boardBad fails column series 
	assert.Equal(t, false, seriesIsLegal(boardBad, 10, 0, 0, 0, 1))

	// boardBad2 passes row series
	assert.Equal(t, true, seriesIsLegal(boardBad2, 10, 0, 0, 1, 0))
	
	// boardBad2 passes column series
	assert.Equal(t, true, seriesIsLegal(boardBad2, 10, 0, 0, 0, 1))

	// boardBad2 fails diagonal series
	assert.Equal(t, false, seriesIsLegal(boardBad2, 10, 0, 0, 1, 1))
	

}