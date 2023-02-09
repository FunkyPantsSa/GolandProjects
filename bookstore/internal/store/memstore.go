// internal/store/memstore.go

package store

import (
	factory "bookstore/internal/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*Book
}
