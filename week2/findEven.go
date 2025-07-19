package main

import (
	"fmt"
)

func FindEvenNumbers(nums []int) ([]int, error) {
	var evens []int
	for _, n := range nums {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	if len(evens) == 0 {
		return []int{}, fmt.Errorf("No even numbers found")
	}
	return evens, nil
}

func main() {
	input1 := []int{1, 3, 2, 4, 5, 6, 7, 8}
	evens1, err1 := FindEvenNumbers(input1)
	fmt.Println(evens1, err1)
	input2 := []int{}
	evens2, err2 := FindEvenNumbers(input2)
	fmt.Println(evens2, err2)
}