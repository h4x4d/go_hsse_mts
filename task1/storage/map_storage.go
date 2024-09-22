package storage

import "task1/book"

type MapStorage struct {
	storage map[uint64]*book.NumberedBook
}

func (storage *MapStorage) GetBook(id uint64) (*book.NumberedBook, bool) {
	foundBook, ok := storage.storage[id]
	return foundBook, ok
}

func (storage *MapStorage) AddBook(toAdd *book.NumberedBook) {
	if storage.storage == nil {
		storage.storage = make(map[uint64]*book.NumberedBook, 1)
	}
	storage.storage[toAdd.Id] = toAdd
}

func (storage *MapStorage) DeleteBook(id uint64) {
	delete(storage.storage, id)
}

func (storage *MapStorage) Clear() {
	storage.storage = make(map[uint64]*book.NumberedBook)
}

func (storage *MapStorage) GetAllBooks() []*book.NumberedBook {
	booksSlice := make([]*book.NumberedBook, 0, len(storage.storage))
	for _, bookPtr := range storage.storage {
		booksSlice = append(booksSlice, bookPtr)
	}
	return booksSlice
}
