package storage

import (
	"context"

	"github.com/mokh1rbek/film_CRUD/models"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
	Country() CategoryRepoI
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *models.CreateCategory) (string, error)
	GetByPKey(ctx context.Context, req *models.CategoryPrimarKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(ctx context.Context, id string, req *models.UpdateCategory) (int64, error)
	Delete(ctx context.Context, req *models.CategoryPrimarKey) error
}

type CountryRepoI interface {
	// Create(ctx context.Context, req *models.CreateCategory) (string, error)
	// GetByPKey(ctx context.Context, req *models.CategoryPrimarKey) (*models.Category, error)
	// GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	// Update(ctx context.Context, id string, req *models.UpdateCategory) (int64, error)
	// Delete(ctx context.Context, req *models.CategoryPrimarKey) error
}
