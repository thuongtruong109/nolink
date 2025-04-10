package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/thuongtruong109/nolink/helpers"
	"github.com/thuongtruong109/nolink/routes"
)

func setupRoutes(app *fiber.App) {
	app.Post("/", routes.ShortenURL)
	app.Post("", routes.ShortenURL)
	app.Get("/:url", routes.ResolveURL)
	app.Post("/own", routes.AllShortedURLsOfUser)
	app.Post("/all", routes.AllShortedURLs)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "02-Sep-2025 15:04:05",
		TimeZone:   "Asia/Ho_Chi_Minh",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		// AllowOrigins: "http://localhost:3000, http://localhost:3001, https://nolink.vercel.app",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodDelete,

		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setupRoutes(app)

	log.Fatal(app.Listen(helpers.APP_PORT))
}
