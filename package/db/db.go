package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// ConnectDB establishes a connection to the database.
func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
	return rdb
}
