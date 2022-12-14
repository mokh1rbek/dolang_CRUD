package models

import "time"

type FilmActor struct {
	ActorId    int       `json:"actor_id"`
	FilmId     int       `json:"film_id"`
	LastUpdate time.Time `json:"last_update"`
}
