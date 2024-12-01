package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"product-service/internal/interfaces/adapter"

	_ "github.com/mattn/go-sqlite3"
	"go.elastic.co/apm/v2"
)

type sqliteClient struct {
	db *sql.DB
}

func NewSqliteClient() (adapter.DatabaseClient, error) {
	db, err := sql.Open("sqlite3", "./products.db")
	if err != nil {
		return nil, err
	}

	// Create the 'product' table
	createProductsTable := `
	CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Execute the table creation queries
	_, err = db.Exec(createProductsTable)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &sqliteClient{
		db: db,
	}, nil
}

func (r *sqliteClient) Execute(ctx context.Context, query string, args ...interface{}) adapter.ExecuteResult {
	span, ctx := apm.StartSpan(ctx, "Execute", "database")
	defer span.End()

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		apm.CaptureError(ctx, err)
		return adapter.ExecuteResult{
			Error: fmt.Errorf("failed to execute query: %w", err),
		}
	}

	lastInsertID, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()

	return adapter.ExecuteResult{
		Result:       result,
		LastInsertID: lastInsertID,
		RowsAffected: rowsAffected,
		Error:        nil,
	}
}

func (r *sqliteClient) QueryRows(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	span, ctx := apm.StartSpan(ctx, "QueryRows", "database")
	defer span.End()

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		apm.CaptureError(ctx, err)
		return nil, err
	}

	return rows, nil
}

func (r *sqliteClient) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	span, ctx := apm.StartSpan(ctx, "QueryRow", "database")
	defer span.End()

	// Execute the query
	row := r.db.QueryRowContext(ctx, query, args...)

	return row
}

func (s *sqliteClient) Close() error {
	err := s.db.Close()

	if err != nil {
		return err
	}

	return nil
}
