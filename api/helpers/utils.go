package helpers

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var (
	APP_PORT = Env("APP_PORT", ":8080")
	DOMAIN = Env("DOMAIN", "localhost:8080")
	DOMAIN_RETURN = Env("DOMAIN_RETURN", "http://localhost:8080")
	API_QUOTA = Env("API_QUOTA", "100")
	DB_ADDR = Env("DB_ADDR", "localhost:6379")

	DB_USER = Env("DB_USER", "root")
	DB_PASS = Env("DB_PASS", "")
)

func Env(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func EnforceHTTP(url string) string {
	if url[:4] != HTTP_PREFIX {
		return HTTP + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == DOMAIN {
		return false
	}

	newURL := strings.Replace(url, HTTP, "", 1)
	newURL = strings.Replace(newURL, HTTPS, "", 1)
	newURL = strings.Replace(newURL, WWW, "", 1)

	newURL = strings.Split(newURL, "/")[0]

	return newURL != DOMAIN
}

func ResponseErr(c *fiber.Ctx, status int, err string) error {
	return c.Status(status).JSON(fiber.Map{ "error": err })
}