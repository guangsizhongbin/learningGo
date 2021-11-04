package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// To connect to a Redis Server
func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

// To connect to a Redis Server using a connection string
func InitRedisWithString() *redis.Client {
	// opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	opt, err := redis.ParseURL("redis://localhost:6379/")
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	return rdb
}

func QuickStartRedis(rdb *redis.Client) {
	// 设置一个key, value
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// 其 ttl 为 -1
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// 获取一个不存在的key
	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func TestRedisWithNX(t *testing.T) {
	rdb := InitRedis()

	// 设置一个10秒过期的key
	b, err := rdb.SetNX(ctx, "key0", "value", 10*time.Second).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

	// 设置其过时间为 -1
	b, err = rdb.SetNX(ctx, "key1", "value", redis.KeepTTL).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}

func TestRedisWithSort(t *testing.T) {
	rdb := InitRedis()
	// SORT list LIMIT 0 2 ASC
	vals, err := rdb.Sort(ctx, "list", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)

}
