package main

import (
	"fmt"
	"live-projects/recursion"
)

func main() {
    var n int64
    for n = 0; n <= 21; n++ {
        fmt.Printf("%3d! = %20d\n", n, recursion.Factorial(n))
    }
    fmt.Println()
}