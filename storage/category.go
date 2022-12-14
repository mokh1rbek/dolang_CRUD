package storage

import (
	"database/sql"

	"github.com/mokh1rbek/golang_CRUD/models"
	
)

func CreateCategory(db *sql.DB, category models.FilmCategory) (int, error) {

	var (
		id    int
		query string
	)

	query = `
		INSERT INTO 
			users (film_id, category_id)
		VALUES ( $1, $2 )
		RETURNING category_id
	`
	err := db.QueryRow(
		query,
		category.FilmId,
		category.CategoryId,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetCategoryById(db *sql.DB, id int) (models.FilmCategory, error) {

	var (
		category models.FilmCategory
		query    string
	)

	query = `
		SELECT
			film_id, 
			last_update
		FROM
			film_category
		WHERE category_id = $1
	`
	err := db.QueryRow(
		query,
		id,
	).Scan(
		&category.CategoryId,
		&category.FilmId,
		&category.LastUpdate,
	)

	if err != nil {
		return models.FilmCategory{}, err
	}

	return category, nil
}

func GetCategoryList(db *sql.DB) ([]models.FilmCategory, error) {

	var (
		categories []models.FilmCategory
		query      string
	)

	query = `
		SELECTid
			category_id,
			film_id,
			last_update
		FROM
			film_category
	`
	rows, err := db.Query(query)

	if err != nil {
		return []models.FilmCategory{}, err
	}

	for rows.Next() {
		var category models.FilmCategory

		err = rows.Scan(
			&category.CategoryId,
			&category.FilmId,
			&category.LastUpdate,
		)

		if err != nil {
			return []models.FilmCategory{}, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func UpdateCategory(db *sql.DB, category models.FilmCategory) (int64, error) {

	var (
		query string
	)

	query = `
		UPDATE
			film_category
		SET
			film_id = $2,
			last_update = $3
		WHERE
			category_id = $1
	`

	result, err := db.Exec(
		query,
		category.CategoryId,
		category.FilmId,
		category.LastUpdate,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteCategory(db *sql.DB, id int) error {

	_, err := db.Exec(`DELETE FROM film_category WHERE category_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
