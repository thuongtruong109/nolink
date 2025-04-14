package helpers

import (
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDR"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})

	return rdb
}


func EnforceHTTP(url string) string {
	if url[:4] != HTTP_PREFIX {
		return HTTP + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	newURL := strings.Replace(url, HTTP, "", 1)
	newURL = strings.Replace(newURL, HTTPS, "", 1)
	newURL = strings.Replace(newURL, WWW, "", 1)

	newURL = strings.Split(newURL, "/")[0]

	return newURL != os.Getenv("DOMAIN")
}

func ResponseErr(c *fiber.Ctx, status int, err string) error {
	return c.Status(status).JSON(fiber.Map{ "error": err })
}