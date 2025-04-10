package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/thuongtruong109/nolink/helpers"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := helpers.CreateClient(0)
	defer r.Close()

	value, err := r.Get(helpers.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": helpers.URL_NOT_FOUND})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": helpers.CANNOT_RETRIEVE_TO_DB})
	}

	rInr := helpers.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(helpers.Ctx, helpers.COUNTER)
	return c.Redirect(value, 301)
}
