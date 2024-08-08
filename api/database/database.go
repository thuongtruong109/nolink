package database

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong109/gouse/connection"
)

func CreateClient(dbNo int) *redis.Client {
	addr := os.Getenv("DB_ADDR")
	pass := os.Getenv("DB_PASS")
	rdb := connection.Redis(addr, pass, dbNo)

	return rdb
}
