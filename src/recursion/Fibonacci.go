package recursion

func Fibonacci(n int64) int64 {
	if n < 0 {
		panic("Unsupported value: n < 0")
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n - 2) + Fibonacci(n - 1)
}

