package redis

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestRedisWihtUnsupportedCommand(t *testing.T) {
	rdb := InitRedisWithString()

	val, err := rdb.Do(ctx, "get", "key").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println(val.(string))
}
