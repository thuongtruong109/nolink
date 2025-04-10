package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/thuongtruong109/nolink/database"
	"github.com/thuongtruong109/nolink/helpers"
)

func AllShortedURLsOfUser(c *fiber.Ctx) error {
	ip := c.Query("ip")

	r := database.CreateClient(0)
	defer r.Close()

	result, err := r.SMembers(helpers.Ctx, helpers.INDEX+ip).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": helpers.CANNOT_RETRIEVE_CLIENT})
	}

	var updatedResult []string

	for _, member := range result {
		value, err := r.Get(helpers.Ctx, member).Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": helpers.CANNOT_RETRIEVE_URLS_OF_CLIENT})
		}

		updatedResult = append(updatedResult, helpers.DOMAIN+"/"+member)

		fmt.Printf("Key: %s, Value: %s\n", member, value)
	}

	return c.Status(fiber.StatusOK).JSON(updatedResult)
}

func AllShortedURLs(c *fiber.Ctx) error {
	r := database.CreateClient(0)
	defer r.Close()

	result, err := r.Keys(helpers.Ctx, "*").Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": helpers.CANNOT_RETRIEVE_ALL_URLS})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"count": len(result)})
}
