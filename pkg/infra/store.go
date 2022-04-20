package infra

import (
	"hex/pkg/types"
)

type Store interface {
	All() ([]types.Product, error)
	Find(string) (*types.Product, error)
	Insert(types.Product) error
	Delete(string) error
}
