// store/factory/factory.go

package factory

import (
	"bookstore/store"
	"errors"
	"fmt"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]store.Store)
)

func Register(name string, p store.Store) error {
	providersMu.Lock()
	defer providersMu.Unlock()
	if p == nil {
		//panic("store: Register provider is nil")

		return errors.New("store: Register provider is nil")
	}

	if _, dup := providers[name]; dup {
		//panic("store: Register called twice for provider " + name
		//fmt.Println("store: Register called twice for provider " + name)
		return errors.New("store: Register called twice for provider " + name)
	}
	providers[name] = p

	return errors.New("err")
}

func New(providerName string) (store.Store, error) {
	providersMu.RLock()
	p, ok := providers[providerName]
	providersMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("store: unknown provider %s", providerName)
	}

	return p, nil
}
