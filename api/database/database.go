package database

import (
	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong109/nolink/helpers"
)

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     helpers.DB_ADDR,
		Username: helpers.DB_USER,
		Password: helpers.DB_PASS,
		DB:       dbNo,
	})

	return rdb
}
