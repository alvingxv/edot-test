package app

import (
	"context"
	"user-service/internal/adapters/database/sqlite"
	"user-service/internal/interfaces/adapter"
)

type Dependencies struct {
	sqlitedb adapter.DatabaseClient
}

func NewDependencies() (*Dependencies, error) {

	db, err := sqlite.NewSqliteClient()
	if err != nil {
		panic(err)
	}

	return &Dependencies{
		sqlitedb: db,
	}, nil
}

func (d *Dependencies) Close(ctx context.Context) error {
	err := d.sqlitedb.Close()
	if err != nil {
		return err
	}

	return nil
}
