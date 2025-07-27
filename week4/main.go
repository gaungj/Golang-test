package main

import (
	"fmt"
	"golang-test/week3/book"
)

func main() {
	bookList := []book.Book{
		{"Golang Basics", 3},
		{"Advanced Programming", 2},
		{"Algorithms and Data Structures", 1},
	}

	book.DisplayBooksInfo(bookList)

	if book.BorrowBookByTitle(bookList, "Golang Basics") {
		fmt.Println("Successfully borrowed book: Golang Basics")
	} else {
		fmt.Println("Book not available: Golang Basics")
	}

	book.ReturnBookByTitle(bookList, "Advanced Programming")
	fmt.Println("Book returned: Advanced Programming")
	
	if bookList[0].Borrow() {
		fmt.Println("Successfully borrowed book:", bookList[0].Title)
	} else {
		fmt.Println("Book not available:", bookList[0].Title)
	}

	bookList[1].Return()
	fmt.Println("Book returned:", bookList[1].Title)

	book.DisplayBooksInfo(bookList)
}