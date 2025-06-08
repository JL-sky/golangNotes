package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func redisHash(ctx context.Context, redisClient *redis.Client) {
	// 单个设置
	redisClient.HSet(ctx, "map", "name", "jack")
	// 批量设置
	redisClient.HMSet(ctx, "map", map[string]interface{}{"a": "b", "c": "d", "e": "f"})
	// 单个访问
	fmt.Println(redisClient.HGet(ctx, "map", "a").Val())
	// 批量访问
	fmt.Println(redisClient.HMGet(ctx, "map", "a", "b").Val())
	// 获取整个map
	fmt.Println(redisClient.HGetAll(ctx, "map").Val())
}
