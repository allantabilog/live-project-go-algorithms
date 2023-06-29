package recursion

var fibonacciValues []int64
// Pre-calculate the first 93 fibonacci numbers
// and store them in a slice
func InitialiseSlice() {
	fibonacciValues = make([]int64, 93) 
	fibonacciValues[0] = 0
	fibonacciValues[1] = 1

	for i := 2; i < 93; i++ {
		fibonacciValues[i] = fibonacciValues[i - 1] + fibonacciValues[i - 2]
	}
}

// no more calculation required 
// basically,just look up the answer 
// that was precalculated and stored in the slice
func FibonacciPrefilled(n int64) int64 {
	return fibonacciValues[n]
}

