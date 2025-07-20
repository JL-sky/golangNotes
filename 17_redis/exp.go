package main

import (
	"fmt"

	"sync"

	"github.com/go-redis/redis"
)

var client *redis.Client
var once sync.Once

func Init() *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

func exp() {
	rdb := Init()
	// 设置键值对, 并设置过期时间为0
	err := rdb.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	// 获取键值对
	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// 批量设置键值对
	pipe := rdb.Pipeline()
	pipe.Set("key1", "value1", 0)
	pipe.Set("key2", "value2", 0)
	_, err = pipe.Exec()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", rdb.Get("key1").Val())
	fmt.Println("key2", rdb.Get("key2").Val())
}
