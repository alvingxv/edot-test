package productuc

import (
	"context"
	"net/http"
	"product-service/internal/interfaces/usecase"
	"product-service/pkg/dto"

	"go.elastic.co/apm/v2"
)

func (uc *productUsecase) GetProducts(ctx context.Context, req *usecase.GetProductsRequest) *dto.Response {
	apmSpan, ctx := apm.StartSpan(ctx, "GetProducts", "usecase")
	defer apmSpan.End()

	resp := dto.New()

	products, err := uc.productRepository.GetProductsFromDb(ctx)
	if err != nil {
		resp.SetError(http.StatusNotFound, err.Status(), err.Message(), err)
		return resp
	}

	resp.SetSuccess(http.StatusOK, "00", "Success retrieve products", products)

	return resp
}
