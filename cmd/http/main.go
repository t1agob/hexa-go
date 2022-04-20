package main

import (
	"hex/pkg/domain"
	"hex/pkg/infra/api"
	"hex/pkg/infra/store"
)

func main() {
	store := store.NewMemoryStore()
	domain := domain.NewProductService(store)
	rest := api.NewRestAPI(domain)

	rest.Run()
}
