package repository

import (
	"context"
	"product-service/pkg/errs"
	"time"
)

type ProductRepository interface {
	GetProductsFromDb(ctx context.Context) ([]Product, errs.MessageErr)
}

type Product struct {
	Id        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
