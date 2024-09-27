package library

import (
	"github.com/stretchr/testify/assert"
	"hash/fnv"
	"task1/book"
	"task1/storage"
	"testing"
	"time"
)

var storages = []storage.Storage{
	&storage.MapStorage{},
	&storage.SliceStorage{},
}

func Hash64(str string) uint64 {
	hash := fnv.New64a()
	_, err := hash.Write([]byte(str))
	if err != nil {
		panic(err)
	}
	return hash.Sum64()
}

func Hash32(str string) uint64 {
	hash := fnv.New32a()
	_, err := hash.Write([]byte(str))
	if err != nil {
		panic(err)
	}
	return uint64(hash.Sum32())
}

var booksLib = []book.Book{
	{
		Title:         "Harry Potter",
		Author:        "J. K. Rowling",
		DatePublished: time.Date(1997, time.June, 26, 0, 0, 0, 0, time.UTC),
		Language:      "en",
		Publisher:     "Bloomsbury Publishing",
	},
	{
		Title:         "The Witcher",
		Author:        "Andrzej Sapkowski",
		DatePublished: time.Date(1986, time.January, 1, 0, 0, 0, 0, time.UTC),
		Language:      "pl",
		Publisher:     "SuperNowa",
	},
	{
		Title:         "War and Peace",
		Author:        "Leo Tolstoy",
		DatePublished: time.Date(1865, time.January, 1, 0, 0, 0, 0, time.UTC),
		Language:      "ru",
		Publisher:     "The Russian Messenger",
	},
}

func TestLibrary(t *testing.T) {
	for index, testStorage := range storages {
		booksCopy := booksLib
		library := NewLibrary(testStorage, Hash64)

		for _, bookToAdd := range booksCopy {
			library.AddBook(bookToAdd)
			addedBook, status := library.GetBook(bookToAdd.Title)
			assert.Equal(t, status, true, "Added book not found")
			assert.Equal(t, addedBook, bookToAdd, "Incorrect book added")
		}

		library.DeleteBook(booksCopy[0].Title)
		_, status := library.GetBook(booksCopy[0].Title)
		assert.Equal(t, status, false, "Book is not deleted correctly")

		bookInLibrary, _ := library.GetBook(booksCopy[1].Title)

		library.RebuildHash(Hash32)

		newHashBook, status := library.GetBook(bookInLibrary.Title)
		assert.Equal(t, bookInLibrary, newHashBook, "Book in Library changed after hash change")

		library.RebuildStorage(storages[1-index])

		newStorageBook, status := library.GetBook(bookInLibrary.Title)
		assert.Equal(t, bookInLibrary, newStorageBook, "Book in Library changed after hash change")

		library.AddBook(booksCopy[0])
		addedBook, status := library.GetBook(booksCopy[0].Title)
		assert.Equal(t, status, true, "Added book not found")
		assert.Equal(t, addedBook, booksCopy[0], "Incorrect book added")
	}
}
