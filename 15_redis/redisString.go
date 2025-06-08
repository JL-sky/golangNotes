package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func redisString(ctx context.Context, redisClient *redis.Client) {
	// 简单存取
	redisClient.Set(ctx, "name", "jack", 0)
	val, err := redisClient.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	// 获取类型
	fmt.Println(redisClient.Type(ctx, "name"))

	// 设置过期时间
	redisClient.Expire(ctx, "name", time.Second*2)
	fmt.Println(redisClient.Get(ctx, "name").Val())
	time.Sleep(time.Second * 3)
	fmt.Println(redisClient.Get(ctx, "name").Val())
	// 获取过期时间
	fmt.Println(redisClient.TTL(ctx, "name"))
	fmt.Println(redisClient.PTTL(ctx, "name"))
	redisClient.Set(ctx, "name", "jack", 2)
	// 取消过期时间
	redisClient.Persist(ctx, "name")
	time.Sleep(time.Second * 2)
	fmt.Println(redisClient.Get(ctx, "name"))

	// 批量存取
	redisClient.MSet(ctx, "cookie", "12345", "token", "abcefg")
	fmt.Println(redisClient.MGet(ctx, "cookie", "token").Val())

	redisClient.Set(ctx, "age", "1", 0)
	// 自增
	redisClient.Incr(ctx, "age")
	fmt.Println(redisClient.Get(ctx, "age").Val())
	// 自减
	redisClient.Decr(ctx, "age")
	fmt.Println(redisClient.Get(ctx, "age").Val())
}
