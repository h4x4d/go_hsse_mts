package library

import (
	"task1/book"
	"task1/storage"
)

type Library struct {
	storage storage.Storage

	hashFunc func(string) uint64
}

func NewLibrary(storage storage.Storage, hashFunc func(string) uint64) *Library {
	return &Library{storage, hashFunc}
}

func (library *Library) RebuildHash(newHash func(string) uint64) {
	save := library.storage.GetAllBooks()
	library.storage.Clear()

	for _, bookPtr := range save {
		bookPtr.Id = newHash(bookPtr.Title)
		library.storage.AddBook(bookPtr)
	}
	library.hashFunc = newHash
}

func (library *Library) RebuildStorage(newStorage storage.Storage) {
	for _, bookPtr := range library.storage.GetAllBooks() {
		newStorage.AddBook(bookPtr)
	}
	library.storage.Clear()
	library.storage = newStorage
}

func (library *Library) GetBook(title string) (book.Book, bool) {
	bookPtr, ok := library.storage.GetBook(library.hashFunc(title))
	return bookPtr.Book, ok
}

func (library *Library) AddBook(bookToAdd book.Book) {
	bookWithId := book.NumberedBook{Id: library.hashFunc(bookToAdd.Title), Book: bookToAdd}
	library.storage.AddBook(bookWithId)
}

func (library *Library) DeleteBook(title string) {
	library.storage.DeleteBook(library.hashFunc(title))
}
