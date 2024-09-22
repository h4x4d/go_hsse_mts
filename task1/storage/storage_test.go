package storage

import (
	"github.com/stretchr/testify/assert"
	"task1/book"
	"testing"
	"time"
)

var storages = []Storage{
	&MapStorage{},
	&SliceStorage{},
}

var books = []book.NumberedBook{
	{
		Id: 1,
		Book: book.Book{
			Title:         "Harry Potter",
			Author:        "J. K. Rowling",
			DatePublished: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
			Language:      "en",
			Publisher:     "Bloomsbury Publishing",
		},
	},
	{
		Id: 121,
		Book: book.Book{
			Title:         "The Witcher",
			Author:        "Andrzej Sapkowski",
			DatePublished: time.Date(1986, time.January, 1, 0, 0, 0, 0, time.UTC),
			Language:      "pl",
			Publisher:     "SuperNowa",
		},
	},
	{
		Id: 2345,
		Book: book.Book{
			Title:         "War and Peace",
			Author:        "Leo Tolstoy",
			DatePublished: time.Date(1865, time.January, 1, 0, 0, 0, 0, time.UTC),
			Language:      "ru",
			Publisher:     "The Russian Messenger",
		},
	},
}

func TestCRUD(t *testing.T) {
	for _, storage := range storages {
		booksCopy := books
		for index, bookToAdd := range booksCopy {
			storage.AddBook(&bookToAdd)
			assert.Equal(t, index+1, len(storage.GetAllBooks()), "Book is not added")
			addedBook, status := storage.GetBook(bookToAdd.Id)
			assert.Equal(t, status, true, "Added book not found")
			assert.Equal(t, *addedBook, bookToAdd, "Incorrect book added")
		}

		storage.DeleteBook(booksCopy[0].Id)
		_, status := storage.GetBook(booksCopy[0].Id)
		assert.Equal(t, status, false, "Book is not deleted correctly")

		bookToEdit, _ := storage.GetBook(booksCopy[1].Id)
		bookToEdit.Language = "en"
		bookInStorage, _ := storage.GetBook(booksCopy[1].Id)
		assert.Equal(t, bookInStorage, bookToEdit, "Book in storage is not modified on modify of book")

		storage.Clear()
		assert.Equal(t, len(storage.GetAllBooks()), 0, "Books not deleted on clear")

	}
}
