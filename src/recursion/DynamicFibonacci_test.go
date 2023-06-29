package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciPrefilled(t *testing.T) {
	type TestData struct {
		input int64
		expected int64
	}
	tests := []TestData{
		{input: int64(0), expected: int64(0)},
		{input: int64(1), expected: int64(1)},
		{input: int64(5), expected: int64(5)},
		{input: int64(10), expected: int64(55)},
		{input: int64(15), expected: int64(610)},
		{input: int64(20), expected: int64(6765)},
		{input: int64(25), expected: int64(75025)},
		{input: int64(30), expected: int64(832040)},
		{input: int64(35), expected: int64(9227465)},
		{input: int64(40), expected: int64(102334155)},
	}

	InitialiseSlice()
	for _, tc := range tests {
		assert.Equal(t, tc.expected, FibonacciPrefilled(tc.input))
	}
}