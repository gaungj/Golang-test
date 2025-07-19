package main

import (
	"fmt"
	"strings"
	"unicode"
)

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

func main() {
	// 2. palindrome
	testCases := []string{
		"radar",
		"hello",
	}

	for _, s := range testCases {
		fmt.Println(isPalindrome(s))
	}
}

