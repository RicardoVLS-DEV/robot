package store

import (
	"context"
	"errors"
	"fmt"
	"robot/internal/domain"

	"github.com/jackc/pgx/v5"
)

type CategoryStore struct {
	store *Store
}

func NewCategoryStore(s *Store) *CategoryStore {
	return &CategoryStore{
		store: s,
	}
}

func (st *CategoryStore) Insert(ctx context.Context, category *domain.Category) (domain.CategoryID, error) {
	query := `
		INSERT INTO category
		(name)
		VALUES ($1)
		RETURNING id
		`

	var id domain.CategoryID
	err := st.store.db.QueryRow(ctx, query, category.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert category: %w", err)
	}

	return id, nil
}

func (st *CategoryStore) FindByName(ctx context.Context, name string) (*domain.Category, error) {
	op := "FindByName"
	
	query := `
		SELECT id, name
		FROM category
		WHERE name = $1
	`
	
	var id domain.CategoryID
	var categoryName string

	err := st.store.db.QueryRow(ctx, query, name).Scan(&id, &categoryName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NewRobotErr(op, "name", name, domain.ErrNotFound, "")
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	category, err := domain.NewCategory(categoryName)
	if err != nil {
		return nil, err
	}

	return category, nil
}
