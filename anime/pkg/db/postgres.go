package db

import (
	"anime/config"
	"anime/models"
	"anime/pkg/parser"
	"context"
	"fmt"
	"log"
	"sync"

	// "log"
	// "net/http"
	// "strconv"
	// "strings"

	// "github.com/PuerkitoBio/goquery"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect() (*PGRepo, error) {
	login, _ := config.Get("DB_LOGIN")
	password, _ := config.Get("DB_PASSWORD")

	URL := fmt.Sprintf("postgres://%s:%s@localhost:5432/anime",
		login,
		password)

	pool, err := pgxpool.Connect(context.Background(), URL)
	if err != nil {
		return nil, err
	}

	return &PGRepo{
		mu:   sync.Mutex{},
		pool: pool,
	}, nil
}

func (repo *PGRepo) FillAnimeTable() error {
	animes := parser.GetAnimesFromSite(10000)

	for _, anime := range animes {
		cmd := "INSERT INTO anime_info (title, description, episodes, type) VALUES($1, $2, $3, $4);"
		q, err := repo.pool.Query(context.Background(),
			cmd,
			anime.Title,
			anime.Desc,
			anime.Episodes,
			anime.Type)
		if err != nil {
			log.Fatal(err)
		}
		q.Close()
	}

	return nil
}

func (repo *PGRepo) GetAnimes() ([]*models.Anime, error) {
	var animes []*models.Anime

	cmd := "SELECT id, title, description, episodes, type from anime_info"
	rows, err := repo.pool.Query(context.Background(),
		cmd)

	if err != nil {
		log.Println("Cannot get animes from database!", err)
		return animes, err
	}
	defer rows.Close()
	for rows.Next() {
		anime := &models.Anime{}
		err = rows.Scan(
			&anime.ID,
			&anime.Title,
			&anime.Desc,
			&anime.Episodes,
			&anime.Type,
		)
		if err != nil {
			log.Println(err)
			return animes, err
		}
		animes = append(animes, anime)
	}

	return animes, nil
}
