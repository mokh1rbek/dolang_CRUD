package storage

import (
	"database/sql"

	"github.com/mokh1rbek/golang_CRUD/models"
)

func CreateFilm(db *sql.DB, film models.Film) (int, error) {

	var (
		id    int
		query string
	)

	query = `
		INSERT INTO 
			film (title, description,release_year,language_id,rental_duration,rental_rate,length,replacement_cost,rating,last_update,special_features,fulltext) 
		VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12 )
		RETURNING film_id
	`
	err := db.QueryRow(
		query,
		film.Title,
		film.Description,
		film.ReleaseYear,
		film.LanguageId,
		film.RentalDuration,
		film.RentalRate,
		film.Length,
		film.ReplacementCost,
		film.Rating,
		film.LastUpdate,
		film.SpecialFeatures,
		film.Fulltext,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetFilmById(db *sql.DB, id int) (models.Film, error) {

	var (
		film  models.Film
		query string
	)

	query = `
		SELECT
			film_id,
			title, 
			description,
			release_year,
			language_id,
			rental_duration,
			rental_rate,
			length,
			replacement_cost,
			rating,
			last_update,
			special_features,
			fulltext
		FROM
			film
		WHERE id = $1
	`
	err := db.QueryRow(
		query,
		id,
	).Scan(
		&film.FilmId,
		&film.Title,
		&film.Description,
		&film.ReleaseYear,
		&film.LanguageId,
		&film.RentalDuration,
		&film.RentalRate,
		&film.Length,
		&film.ReplacementCost,
		&film.Rating,
		&film.LastUpdate,
		&film.SpecialFeatures,
		&film.Fulltext,
	)

	if err != nil {
		return models.Film{}, err
	}

	return film, nil
}

func GetFilmList(db *sql.DB) ([]models.Film, error) {

	var (
		film  []models.Film
		query string
	)

	query = `
		SELECT
			film_id,
			title, 
			description,
			release_year,
			language_id,
			rental_duration,
			rental_rate,
			length,
			replacement_cost,
			rating,
			last_update,
			special_features,
			fulltext
		FROM
				film
	`
	rows, err := db.Query(query)

	if err != nil {
		return []models.Film{}, err
	}

	for rows.Next() {
		var films models.Film

		err = rows.Scan(
			&films.FilmId,
			&films.Title,
			&films.Description,
			&films.ReleaseYear,
			&films.LanguageId,
			&films.RentalDuration,
			&films.RentalRate,
			&films.Length,
			&films.ReplacementCost,
			&films.Rating,
			&films.LastUpdate,
			&films.SpecialFeatures,
			&films.Fulltext,
		)

		if err != nil {
			return []models.Film{}, err
		}

		film = append(film, films)
	}

	return film, nil
}

func UpdateFilm(db *sql.DB, film models.Film) (int64, error) {

	var (
		query string
	)

	query = `
		UPDATE
			film
		SET
			title = $2, 
			description = $3,
			release_year = $4,
			language_id = $5,
			rental_duration = $6,
			rental_rate = $7,
			length = $8,
			replacement_cost = $9,
			rating = $10,
			last_update = $11,
			special_features = $12,
			fulltext = $13 
		WHERE
		film_id = $1
	`

	result, err := db.Exec(
		query,
		film.FilmId,
		film.Title,
		film.Description,
		film.ReleaseYear,
		film.LanguageId,
		film.RentalDuration,
		film.RentalRate,
		film.Length,
		film.ReplacementCost,
		film.Rating,
		film.LastUpdate,
		film.SpecialFeatures,
		film.Fulltext,
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

func DeleteFilm(db *sql.DB, id int) error {

	_, err := db.Exec(`DELETE FROM film WHERE film_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
