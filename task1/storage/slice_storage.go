package storage

import (
	"slices"
	"task1/book"
)

type SliceStorage struct {
	storage []*book.NumberedBook
}

func (storage *SliceStorage) GetBook(id uint64) (*book.NumberedBook, bool) {
	for _, bookPtr := range storage.storage {
		if bookPtr.Id == id {
			return bookPtr, true
		}
	}
	return nil, false
}

func (storage *SliceStorage) DeleteBook(id uint64) {
	for index, bookPtr := range storage.storage {
		if bookPtr.Id == id {
			storage.storage = slices.Delete(storage.storage, index, index+1)
			return
		}
	}
}

func (storage *SliceStorage) Clear() {
	storage.storage = make([]*book.NumberedBook, 0)
}

func (storage *SliceStorage) AddBook(toAdd *book.NumberedBook) {
	storage.storage = append(storage.storage, toAdd)
}

func (storage *SliceStorage) GetAllBooks() []*book.NumberedBook {
	return storage.storage
}
