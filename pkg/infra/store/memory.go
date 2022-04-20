package store

import (
	"errors"
	"hex/pkg/types"
)

type MemoryStore struct {
	store map[string]types.Product
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		store: make(map[string]types.Product),
	}
}

func (m *MemoryStore) All() ([]types.Product, error) {
	productRange := []types.Product{}

	for _, v := range m.store {
		productRange = append(productRange, v)
	}

	return productRange, nil
}

func (m *MemoryStore) Find(id string) (*types.Product, error) {
	p, ok := m.store[id]
	if !ok {
		return nil, errors.New("Item does not exist")
	}

	return &p, nil
}

func (m *MemoryStore) Insert(p types.Product) error {
	m.store[p.Id] = p

	return nil
}

func (m *MemoryStore) Delete(id string) error {
	delete(m.store, id)

	return nil
}
