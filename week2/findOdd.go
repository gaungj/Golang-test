package main

import (
	"errors"
	"fmt"
)

func findOdd(nums []int) ([]int, error) {
	var odds []int
	for _, n := range nums {
		if n%2 != 0 {
			odds = append(odds, n)
		}
	}
	if len(odds) == 0 {
		return nil, errors.New("No odd numbers found")
	}
	return odds, nil
}

func main() {
	input := []int{}
	result, err := findOdd(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	input = []int{2, 4, 7, 9, 10, 12}
	result, err = findOdd(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}