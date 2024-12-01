package productuc

import (
	"product-service/internal/interfaces/repository"
	"product-service/internal/interfaces/usecase"
)

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) usecase.ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}
