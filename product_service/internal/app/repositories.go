package app

import (
	"product-service/internal/interfaces/repository"
	productrepo "product-service/internal/repository/product"
)

type Repositories struct {
	productRepository repository.ProductRepository
}

func NewRepos(dependencies *Dependencies) *Repositories {
	return &Repositories{
		productRepository: productrepo.NewProductRepository(dependencies.sqlitedb),
	}
}
