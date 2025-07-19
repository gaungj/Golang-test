package main

import (
	"fmt"
)

func findMax(nums []int) int {
	if len(nums) == 0 {
		panic("slice is empty")
	}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	fmt.Println(findMax([]int{1, 3, 2, 9, 5}))
	fmt.Println(findMax([]int{}))
}