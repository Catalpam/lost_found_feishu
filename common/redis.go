package common

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println("Redis 连接失败：", pong, err)
		panic("redis ping error")
		return
	}
	fmt.Println("Redis 连接成功：", pong)
}
