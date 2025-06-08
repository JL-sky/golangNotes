package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func subscribe(ctx context.Context, redisClient *redis.Client) {
	// 发送消息到指定频道
	redisClient.Publish(ctx, "channel", "message")
	// 订阅指定频道
	redisClient.Subscribe(ctx, "channel")
	// 查看订阅状态
	redisClient.PubSubNumSub(ctx, "channel")
}
