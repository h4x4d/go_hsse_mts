package storage

import "task1/book"

type Storage interface {
	GetBook(id uint64) (*book.NumberedBook, bool)
	DeleteBook(id uint64)
	Clear()
	AddBook(book *book.NumberedBook)
	GetAllBooks() []*book.NumberedBook
}
