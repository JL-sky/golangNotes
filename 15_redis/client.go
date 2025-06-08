package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func GetClient() (context.Context, *redis.Client) {
	ctx := context.Background()

	// 创建Redis连接客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Password: "123456",
		DB: 0, // 使用默认DB
	})
	return ctx, redisClient
}
