package main

import (
	"fmt"
)

type Book struct {
	Title    string
	Quantity int
}

func (b *Book) Borrow() bool {
	if b.Quantity > 0 {
		b.Quantity--
		return true
	}
	return false
}

func (b *Book) Return() {
	b.Quantity++
}

func BorrowBookByTitle(bookList []Book, title string) bool {
	for i := range bookList {
		if bookList[i].Title == title {
			return bookList[i].Borrow()
		}
	}
	return false
}

func ReturnBookByTitle(bookList []Book, title string) bool {
	for i := range bookList {
		if bookList[i].Title == title {
			bookList[i].Return()
			return true
		}
	}
	return false
}

func DisplayBooksInfo(bookList []Book) {
	fmt.Println("Book List:")
	for _, book := range bookList {
		fmt.Printf("Title: %s, Quantity: %d\n", book.Title, book.Quantity)
	}
}

func main() {
	bookList := []Book{
		{"Golang Basics", 3},
		{"Advanced Programming", 2},
		{"Algorithms and Data Structures", 1},
	}

	DisplayBooksInfo(bookList)

	if BorrowBookByTitle(bookList, "Golang Basics") {
		fmt.Println("Successfully borrowed book: Golang Basics")
	} else {
		fmt.Println("Book not available: Golang Basics")
	}

	ReturnBookByTitle(bookList, "Advanced Programming")
	fmt.Println("Book returned: Advanced Programming")
	
	if bookList[0].Borrow() {
		fmt.Println("Successfully borrowed book:", bookList[0].Title)
	} else {
		fmt.Println("Book not available:", bookList[0].Title)
	}

	bookList[1].Return()
	fmt.Println("Book returned:", bookList[1].Title)

	DisplayBooksInfo(bookList)
}