package storage_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/mokh1rbek/film_CRUD/config"
	"github.com/mokh1rbek/film_CRUD/storage/postgres"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	filmRepo     *postgres.FilmRepo
	categoryRepo *postgres.CategoryRepo
	actorRepo    *postgres.ActorRepo
)

func TestMain(m *testing.M) {

	cfg := config.Load()

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	filmRepo = postgres.NewFilmRepo(pool)

	os.Exit(m.Run())
}
