package main

import (
	"fmt"
)

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	// 1. factorial
	num := 5
	fmt.Println(factorial(num))
}

