package api

import (
	"anime/models"
	"anime/pkg/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var Animes []*models.Anime

func getDB(repo *db.PGRepo) *db.PGRepo {
	return repo
}

func MainPage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello from main page!",
	})
}

func Profile(c *fiber.Ctx) error {
	return c.Render("profile", fiber.Map{
		"name":  "rume",
		"email": "rume@gmail.com",
	})
}

func GetAnimes(c *fiber.Ctx) error {
	return c.Render("base", fiber.Map{
		"animes": Animes})
}

func GetAnimeByID(c *fiber.Ctx) error {
	animeID, err := strconv.Atoi(c.Params("num"))
	if err != nil {
		return err
	}
	var anime *models.Anime
	for _, a := range Animes {
		if a.ID == uint64(animeID) {
			anime = a
		}
	}
	return c.Render("anime", fiber.Map{
		"Title":    anime.Title,
		"Desc":     anime.Desc,
		"Episodes": anime.Episodes,
		"Type":     anime.Type,
	})
}

func signup(c *fiber.Ctx) {

}
