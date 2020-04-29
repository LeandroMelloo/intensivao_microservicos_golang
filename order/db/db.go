package db

import (
	"github.com/go-redis/redis/v7"
)

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}
