package book

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