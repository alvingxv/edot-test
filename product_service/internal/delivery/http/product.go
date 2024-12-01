package http

import (
	"fmt"
	"product-service/internal/interfaces/delivery"
	"product-service/internal/interfaces/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm"
)

type productHandler struct {
	productUsecase usecase.ProductUsecase
	validate       *validator.Validate
}

func NewProdcutHandler(productUsecase usecase.ProductUsecase, validate *validator.Validate) delivery.ProductHandler {
	return &productHandler{
		productUsecase: productUsecase,
		validate:       validate,
	}
}

func (h *productHandler) GetProducts(c *fiber.Ctx) error {
	apmSpan, ctx := apm.StartSpan(c.Context(), "Register", "Handler")
	defer apmSpan.End()

	resp := h.productUsecase.GetProducts(ctx, nil)
	fmt.Println(resp)

	c.Status(resp.HttpCode)
	c.JSON(resp)
	return nil
}
