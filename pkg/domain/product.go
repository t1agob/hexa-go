package domain

import (
	"hex/pkg/infra"
	"hex/pkg/types"
)

type ProductService interface {
	All() ([]types.Product, error)
	Find(string) (*types.Product, error)
	Add(types.Product) error
	Delete(string) error
}

type Products struct {
	store infra.Store
}

func NewProductService(store infra.Store) *Products {
	return &Products{
		store: store,
	}
}

func (s *Products) All() ([]types.Product, error) {
	pp, err := s.store.All()

	if err != nil {
		return nil, err
	}

	return pp, nil
}

func (s *Products) Find(p string) (*types.Product, error) {
	pr, err := s.store.Find(p)

	if err != nil {
		return nil, err
	}

	return pr, nil
}

func (s *Products) Add(p types.Product) error {
	err := s.store.Insert(p)

	if err != nil {
		return err
	}

	return nil
}

func (s *Products) Delete(p string) error {
	err := s.store.Delete(p)

	if err != nil {
		return err
	}

	return nil
}
