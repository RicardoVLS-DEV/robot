package store

import (
	"context"
	_ "embed"
)

//go:embed migration/schema.sql
var schema string

func (st *Store) Migrate(ctx context.Context) error {
	_, err := st.db.Exec(ctx, schema)
	return err
}