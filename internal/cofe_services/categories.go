package cofe_services

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type Category struct {
	ID        uuid.UUID `json:"guid"`
	Title     string    `json:"title"`
	Query     string    `json:"query"`
	Icon      string    `json:"icon"`
	ItemOrder int32
}

type CategoryStore interface {
	GetCategories(ctx context.Context) ([]Category, error)
	// AddCategory(ctx context.Context, c Category) error
	// EditCategory(ctx context.Context, c Category) error
	// DeleteCategory(ctx context.Context, c Category) error
}

type CategoryRepo struct {
	store CategoryStore
}

func NewCategories(store CategoryStore) *CategoryRepo {
	return &CategoryRepo{
		store: store,
	}
}

func (s *CategoryRepo) GetCategories(ctx context.Context) ([]Category, error) {
	items, err := s.store.GetCategories(ctx)

	if err != nil {
		return nil, err
	}
	return items, nil
}
