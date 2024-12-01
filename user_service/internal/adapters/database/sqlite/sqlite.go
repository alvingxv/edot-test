package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"user-service/internal/interfaces/adapter"

	_ "github.com/mattn/go-sqlite3"
	"go.elastic.co/apm/v2"
)

type sqliteClient struct {
	db *sql.DB
}

func NewSqliteClient() (adapter.DatabaseClient, error) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, err
	}

	// Create the 'users' table
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	// Execute the table creation queries
	_, err = db.Exec(createUsersTable)
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

// // Query executes a query that returns multiple rows
// func (r *Repository) Query(ctx context.Context, query string, args ...interface{}) QueryResult {
// 	// Start a new span for tracing
// 	span, ctx := apm.StartSpan(ctx, "Query", "database")
// 	defer span.End()

// 	// Add query to span
// 	span.SetLabel("query", query)

// 	// Execute the query
// 	rows, err := r.db.QueryContext(ctx, query, args...)
// 	if err != nil {
// 		apm.CaptureError(ctx, err)
// 		return QueryResult{
// 			Error: fmt.Errorf("failed to execute query: %w", err),
// 		}
// 	}
// 	defer rows.Close()

// 	// Collect rows into a slice of maps
// 	var result []map[string]interface{}

// 	// Get column names
// 	columns, err := rows.Columns()
// 	if err != nil {
// 		apm.CaptureError(ctx, err)
// 		return QueryResult{
// 			Error: fmt.Errorf("failed to get columns: %w", err),
// 		}
// 	}

// 	for rows.Next() {
// 		// Create a map to hold the row data
// 		rowMap := make(map[string]interface{})

// 		// Create a slice of interfaces to scan into
// 		scanArgs := make([]interface{}, len(columns))
// 		scanDestinations := make([]interface{}, len(columns))

// 		for i := range columns {
// 			scanDestinations[i] = &scanArgs[i]
// 		}

// 		// Scan the row
// 		if err := rows.Scan(scanDestinations...); err != nil {
// 			apm.CaptureError(ctx, err)
// 			return QueryResult{
// 				Error: fmt.Errorf("failed to scan row: %w", err),
// 			}
// 		}

// 		// Populate the map
// 		for i, colName := range columns {
// 			rowMap[colName] = scanArgs[i]
// 		}

// 		// Append to result slice
// 		result = append(result, rowMap)
// 	}

// 	// Check for any errors encountered during iteration
// 	if err := rows.Err(); err != nil {
// 		apm.CaptureError(ctx, err)
// 		return QueryResult{
// 			Error: fmt.Errorf("row iteration error: %w", err),
// 		}
// 	}

// 	return QueryResult{
// 		Rows:  result,
// 		Error: nil,
// 	}
// }

func (s *sqliteClient) Close() error {
	err := s.db.Close()

	if err != nil {
		return err
	}

	return nil
}
