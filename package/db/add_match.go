package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// ConnectDB establishes a connection to the database.
func Add(c *redis.Client, key string, time float64, member int) {
	ctx := context.Background()

	_,err := c.ZAdd(ctx, key, &redis.Z{Score: time, Member: member}).Result()
	if err != nil {
		panic(err)
	}

}


func Get(c *redis.Client, key string, start, stop int64) []string {
	ctx := context.Background()

	res, err := c.ZRange(ctx, key, start, stop).Result()
	if err != nil {
		panic(err)
	}

	return res
}
