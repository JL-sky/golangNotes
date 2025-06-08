package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func redisList(ctx context.Context, redisClient *redis.Client) {
	// 左边添加元素
	err := redisClient.LPush(ctx, "list", "a", "b", "c", "d", "e").Err()
	if err != nil {
		log.Fatalf("LPush 失败: %v", err)
	}

	// 右边添加元素
	err = redisClient.RPush(ctx, "list", "g", "i", "a").Err()
	if err != nil {
		log.Fatalf("RPush 失败: %v", err)
	}

	// 在参考值前面插入值
	err = redisClient.LInsertBefore(ctx, "list", "a", "aa").Err()
	if err != nil {
		log.Fatalf("LInsertBefore 失败: %v", err)
	}

	// 在参考值后面插入值
	err = redisClient.LInsertAfter(ctx, "list", "a", "gg").Err()
	if err != nil {
		log.Fatalf("LInsertAfter 失败: %v", err)
	}

	// 设置指定下标的元素的值
	err = redisClient.LSet(ctx, "list", 0, "head").Err()
	if err != nil {
		log.Fatalf("LSet 失败: %v", err)
	}

	// 获取列表长度
	len, err := redisClient.LLen(ctx, "list").Result()
	if err != nil {
		log.Fatalf("LLen 失败: %v", err)
	}
	fmt.Printf("列表长度: %d\n", len)

	// 左边弹出元素
	leftElem, err := redisClient.LPop(ctx, "list").Result()
	if err != nil {
		log.Fatalf("LPop 失败: %v", err)
	}
	fmt.Printf("左边弹出元素: %s\n", leftElem)

	// 右边弹出元素
	rightElem, err := redisClient.RPop(ctx, "list").Result()
	if err != nil {
		log.Fatalf("RPop 失败: %v", err)
	}
	fmt.Printf("右边弹出元素: %s\n", rightElem)

	// 访问指定下标的元素
	elem, err := redisClient.LIndex(ctx, "list", 1).Result()
	if err != nil {
		log.Fatalf("LIndex 失败: %v", err)
	}
	fmt.Printf("下标1的元素: %s\n", elem)

	// 访问指定范围内的元素
	elems, err := redisClient.LRange(ctx, "list", 0, 3).Result()
	if err != nil {
		log.Fatalf("LRange 失败: %v", err)
	}
	fmt.Printf("范围0-3的元素: %v\n", elems)

	// 删除指定元素（删除所有a）
	count, err := redisClient.LRem(ctx, "list", 0, "a").Result()
	if err != nil {
		log.Fatalf("LRem 失败: %v", err)
	}
	fmt.Printf("删除了 %d 个 a\n", count)

	// 保留指定范围的元素（只保留前两个元素）
	err = redisClient.LTrim(ctx, "list", 0, 1).Err()
	if err != nil {
		log.Fatalf("LTrim 失败: %v", err)
	}

	// 最终结果
	finalElems, err := redisClient.LRange(ctx, "list", 0, -1).Result()
	if err != nil {
		log.Fatalf("LRange 失败: %v", err)
	}
	fmt.Printf("最终列表元素: %v\n", finalElems)
}
