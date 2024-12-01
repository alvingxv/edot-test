package productrepo

import (
	"context"
	"product-service/internal/interfaces/adapter"
	"product-service/internal/interfaces/repository"
	"product-service/pkg/errs"

	"go.elastic.co/apm"
)

type productRepository struct {
	database adapter.DatabaseClient
}

func NewProductRepository(database adapter.DatabaseClient) repository.ProductRepository {
	return &productRepository{
		database: database,
	}
}

func (rp *productRepository) GetProductsFromDb(ctx context.Context) ([]repository.Product, errs.MessageErr) {
	apmSpan, ctx := apm.StartSpan(ctx, "GetProductsFromDb", "Repository")
	defer apmSpan.End()

	query := `select * FROM products;`

	var products []repository.Product

	rows, err := rp.database.QueryRows(ctx, query)
	if err != nil {
		return nil, errs.NewCustomErrs(
			"Failed Insert Database",
			"FD",
			err.Error(),
		)
	}

	for rows.Next() {
		var product repository.Product

		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.CreatedAt)
		if err != nil {
			return nil, errs.NewCustomErrs(
				"Failed Scan",
				"FD",
				err.Error(),
			)
		}

		products = append(products, product)

	}

	return products, nil
}
