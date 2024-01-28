package main

import (
	"anime/pkg/api"
	"anime/pkg/router"
	"net/http"

	"anime/pkg/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"log"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	api.Animes, _ = conn.GetAnimes()

	viewsEngine := html.NewFileSystem(http.Dir("./public/templates"), ".html") // Используйте свой путь к файлам шаблонов
	app := fiber.New(fiber.Config{
		Views: viewsEngine,
	})
	router.SetupRouters(app)
	app.Listen(":8080")

}
