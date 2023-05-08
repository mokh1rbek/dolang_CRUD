package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mokh1rbek/film_CRUD/models"
)

type CountryRepo struct {
	db *pgxpool.Pool
}

func NewCountryRepo(db *pgxpool.Pool) *CountryRepo {
	return &CountryRepo{
		db: db,
	}
}

func (f *CountryRepo) Create(ctx context.Context, category *models.Countries) (string, error) {

	var (
		id = uuid.New().String()
		// query    string
		// nulluuid sql.NullString
	)

	// query = `
	// 	INSERT INTO category(
	// 		category_id,
	// 		name,
	// 		parent_uuid,
	// 		updated_at
	// 	) VALUES ( $1, $2, $3, now())
	// `

	// if category.ParentUUID == "" {
	// 	_, err := f.db.Exec(ctx, query,
	// 		id,
	// 		category.Name,
	// 		nulluuid,
	// 	)

	// 	if err != nil {
	// 		return "", err
	// 	}
	// } else {

	// 	_, err := f.db.Exec(ctx, query,
	// 		id,
	// 		category.Name,
	// 		category.ParentUUID,
	// 	)

	// 	if err != nil {
	// 		return "", err
	// 	}
	// }

	return id, nil
}
