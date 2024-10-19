package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"orca/conf"
)

var Redis *redis.Client

func init() {
	var c = context.Background()
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.GetString("redis.host"), conf.GetString("redis.port")),
		Password: conf.GetString("redis.password"),
		DB:       conf.GetInt("redis.db"),
	})

	_, err := Redis.Ping(c).Result()
	if err != nil {
		panic("Failed to connect redis server")
	}
}
