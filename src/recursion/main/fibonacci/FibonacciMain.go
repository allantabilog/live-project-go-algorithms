package main

import (
	"fmt"
	"live-projects/recursion"
	"strconv"
)

// Rod's main driver code
func main() {
	for {
		var n_string string 
		fmt.Printf("N: ")
		fmt.Scanln(&n_string);

		if len(n_string) == 0 {
			break;
		}

		n, _ := strconv.ParseInt(n_string, 10, 64);
		fmt.Printf("fibonacci(%d) = %d\n", n, recursion.Fibonacci(n))
	}
}