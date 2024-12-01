package usecase

import (
	"context"
	"product-service/pkg/dto"
)

type ProductUsecase interface {
	// Login(ctx context.Context, req *LogKinRequest) *dto.Response
	GetProducts(ctx context.Context, req *GetProductsRequest) *dto.Response
}

type GetProductsRequest struct {
}
