// internal/store/memstore.go

package store

import (
	"bookstore/store"
	"bookstore/store/factory"
	"errors"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*store.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*store.Book
}

func (m MemStore) Create(book *store.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Update(book *store.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) Get(s string) (store.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemStore) GetAll() ([]store.Book, error) {
	//TODO implement me
	panic("implement me")

}

func (m MemStore) Delete(s string) error {
	//TODO implement me
	//panic("implement me")
	return errors.New("1")
}
