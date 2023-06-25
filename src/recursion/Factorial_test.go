package recursion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorialNegative(t *testing.T){
	assert.Panics(t, func() {
		Factorial(-1)
	})
}
func TestFactorial(t *testing.T) {

	//table-driven test
	//https://dave.cheney.net/2019/05/07/prefer-table-driven-tests

	type TestData struct {
		input int64
		expected int64
	}

	tests := []TestData{
		{input: int64(1), expected: int64(1)},
		{input: int64(2), expected: int64(2)},
		{input: int64(3), expected: int64(6)},
		{input: int64(4), expected: int64(24)},
		{input: int64(5), expected: int64(120)},
		{input: int64(6), expected: int64(720)},
		{input: int64(7), expected: int64(5040)},
		{input: int64(8), expected: int64(40320)},
		{input: int64(9), expected: int64(362880)},
		{input: int64(10), expected: int64(3628800)},
		{input: int64(11), expected: int64(39916800)},
		{input: int64(12), expected: int64(479001600)},
		{input: int64(13), expected: int64(6227020800)},
		{input: int64(14), expected: int64(87178291200)},
		{input: int64(15), expected: int64(1307674368000)},
		{input: int64(16), expected: int64(20922789888000)},
		{input: int64(17), expected: int64(355687428096000)},
		{input: int64(18), expected: int64(6402373705728000)},
		{input: int64(19), expected: int64(121645100408832000)},
		{input: int64(20), expected: int64(2432902008176640000)},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.expected, Factorial(tc.input))
	}

}

