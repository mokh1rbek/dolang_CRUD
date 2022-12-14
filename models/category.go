package models

import "time"

type FilmCategory struct {
	FilmId     int       `json:"film_id"`
	CategoryId int       `json:"category_id"`
	LastUpdate time.Time `json:"last_update"`
}
