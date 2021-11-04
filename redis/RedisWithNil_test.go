package redis

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestRedisNil(t *testing.T) {
	rdb := InitRedisWithString()

	val, err := rdb.Get(ctx, "key").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case val == "":
		fmt.Println("value is empty")
	}
}
