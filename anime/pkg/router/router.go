package router

import (
	"anime/pkg/api"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	animes := app.Group("/animes")

	animes.Get("/", api.MainPage)
	animes.Get("/profile", api.Profile)

	base := animes.Group("/base")
	base.Get("/", api.GetAnimes)
	base.Get("/:num", api.GetAnimeByID)
}
