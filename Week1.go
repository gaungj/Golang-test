package main

import (
	"fmt"
	"strings"
	"unicode"
)

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func isPalindrome(s string) bool {
	var cleaned strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned.WriteRune(unicode.ToLower(char))
		}
	}
	cleanStr := cleaned.String()

	left := 0
	right := len(cleanStr) - 1
	for left < right {
		if cleanStr[left] != cleanStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

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
	// 1. factorial
	num := 5
	fmt.Println(factorial(num))

	// 2. palindrome
	testCases := []string{
		"radar",
		"hello",
	}

	for _, s := range testCases {
		fmt.Println(isPalindrome(s))
	}

	// 3. highest number in slice
	numbers := []int{3, 5, 2, 8, 1}
	fmt.Println(findMax(numbers))
}

