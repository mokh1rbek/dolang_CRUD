package storage

import (
	"database/sql"

	"github.com/mokh1rbek/golang_CRUD/models"

)

func CreateActor(db *sql.DB, actor models.FilmActor) (int, error) {

	var (
		id    int
		query string
	)

	query = `
		INSERT INTO 
			film_actor (actor_id, film_id)
		VALUES ( $1, $2 )
		RETURNING actor_id
	`
	err := db.QueryRow(
		query,
		actor.ActorId,
		actor.FilmId,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetActorById(db *sql.DB, id int) (models.FilmActor, error) {

	var (
		actor models.FilmActor
		query string
	)

	query = `
		SELECT
			actor_id,
			film_id,
			last_update
		FROM
			film_actor
		WHERE actor_id = $1
	`
	err := db.QueryRow(
		query,
		id,
	).Scan(
		&actor.ActorId,
		&actor.FilmId,
		&actor.LastUpdate,
	)

	if err != nil {
		return models.FilmActor{}, err
	}

	return actor, nil
}

func GetActorList(db *sql.DB) ([]models.FilmActor, error) {

	var (
		actors []models.FilmActor
		query  string
	)

	query = `
		SELECT
			actor_id,
			film_id,
			last_update
		FROM
			actors
	`
	rows, err := db.Query(query)

	if err != nil {
		return []models.FilmActor{}, err
	}

	for rows.Next() {
		var actor models.FilmActor

		err = rows.Scan(
			&actor.ActorId,
			&actor.FilmId,
			&actor.LastUpdate,
		)

		if err != nil {
			return []models.FilmActor{}, err
		}

		actors = append(actors, actor)
	}

	return actors, nil
}

func UpdateActor(db *sql.DB, actor models.FilmActor) (int64, error) {

	var (
		query string
	)

	query = `
		UPDATE
			film_actor
		SET
			film_id = $2,
			last_update = $3
		WHERE
			actor_id = $1
	`

	result, err := db.Exec(
		query,
		actor.ActorId,
		actor.FilmId,
		actor.LastUpdate,
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

func DeleteActor(db *sql.DB, id int) error {

	_, err := db.Exec(`DELETE FROM film_actor WHERE film_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
