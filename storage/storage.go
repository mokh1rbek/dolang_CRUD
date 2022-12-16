package storage

import (
	"context"

	"github.com/mokh1rbek/CRUD/models"
)

type StorageI interface {
	CloseDB()
	Film() FilmRepoI
	Category() CategoryRepoI
}

type FilmRepoI interface {
	Create(ctx context.Context, req *models.CreateFilm) (string, error)
	GetByPKey(ctx context.Context, req *models.FilmPrimarKey) (*models.Film, error)
	GetList(ctx context.Context, req *models.GetListFilmRequest) (*models.GetListFilmResponse, error)
	Update(ctx context.Context, id string, req *models.UpdateFilm) (int64, error)
	Delete(ctx context.Context, req *models.FilmPrimarKey) error
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *models.CreateCategory) (string, error)
	GetByPKey(ctx context.Context, req *models.CategoryPrimarKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(ctx context.Context, id string, req *models.UpdateCategory) (int64, error)
	Delete(ctx context.Context, req *models.CategoryPrimarKey) error
}
