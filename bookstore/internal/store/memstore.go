// internal/store/memstore.go

package store

import (
	"bookstore/store"
	"bookstore/store/factory"
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
