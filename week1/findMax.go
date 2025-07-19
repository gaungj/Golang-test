package main

import "fmt"

func findMax(nums []int) int {
	if len(nums) == 0 {
		return 0
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
	// 3. highest number in slice
	numbers := []int{3, 5, 2, 8, 1}
	fmt.Println(findMax(numbers))
}

