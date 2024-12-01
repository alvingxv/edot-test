package app

import (
	"product-service/internal/interfaces/usecase"
	productuc "product-service/internal/usecase/product"
)

type Usecases struct {
	ProductUsecase usecase.ProductUsecase
}

func NewUsecases(repos *Repositories) *Usecases {
	return &Usecases{
		ProductUsecase: productuc.NewProductUsecase(repos.productRepository),
	}
}
