package main

import "fmt"

func Swap(a, b *string) {
	*a, *b = *b, *a
}

func main() {
	str1 := "Hello"
	str2 := "World"
	Swap(&str1, &str2)
	fmt.Println(str1, str2)
}