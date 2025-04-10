package routes

import "github.com/gofiber/fiber/v2"

func Combine(app *fiber.App) {
	app.Post("/", ShortenURL)
	app.Post("", ShortenURL)
	app.Get("/:url", ResolveURL)
	app.Post("/own", AllShortedURLsOfUser)
	app.Post("/all", AllShortedURLs)
}
