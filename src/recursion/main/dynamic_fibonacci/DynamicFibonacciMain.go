package main

import (
	"fmt"
	"live-projects/recursion"
)

func main() {
	recursion.InitialiseSlice()
	fmt.Println(recursion.FibonacciPrefilled(40))
	var i int64
	for i=0; i<93; i++ {
		fmt.Println(recursion.FibonacciPrefilled(i))
	}
}