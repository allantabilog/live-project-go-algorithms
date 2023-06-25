package recursion

/**
 * Recursive factorial function
 * Can only calculate correctly up to 20!
 **/
func Factorial(n int64) int64 {
	if n < 0 || n > 20 {
		panic("Unsupported values of n (< 0 or > 20)")
	}
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
