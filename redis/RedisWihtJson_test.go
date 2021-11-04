package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestRedisWithJson(t *testing.T) {
	// client := InitRedis()
	client := InitRedisWithString()

	var ctx = context.Background()

	json, err := json.Marshal(Author{Name: "feng", Age: 22})
	if err != nil {
		panic(err)
	}

	err = client.Set(ctx, "id1234", json, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

}
