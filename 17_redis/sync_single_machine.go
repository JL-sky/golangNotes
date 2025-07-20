package main

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

type redisClient struct {
	rdb *redis.Client
}

func NewClient() *redisClient {
	return &redisClient{
		rdb: Init(),
	}
}

// 模拟并发问题
func (r *redisClient) syncProblem() {
	r.rdb.Set("cnt", 0, 0)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			val, err := r.rdb.Get("cnt").Int()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = r.rdb.Set("cnt", val+1, 0).Err()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(val + 1)
		}()
	}
	wg.Wait()
	finalVal, err := r.rdb.Get("cnt").Int()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Final value:", finalVal)
}

// 使用互斥锁解决并发问题
func (r *redisClient) method1() {
	r.rdb.Set("cnt", 0, 0)

	var wg sync.WaitGroup
	// 互斥锁
	var mu sync.Mutex
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 加锁, 保证操作的原子性
			mu.Lock()
			defer mu.Unlock()
			val, err := r.rdb.Get("cnt").Int()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = r.rdb.Set("cnt", val+1, 0).Err()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(val + 1)
		}()
	}
	wg.Wait()
	finalVal, err := r.rdb.Get("cnt").Int()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Final value:", finalVal)
}

// 使用Redis的INCR命令
func (r *redisClient) method2() {
	r.rdb.Set("cnt", 0, 0)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// INCR命令
			val, err := r.rdb.Incr("cnt").Result()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(val)
		}()
	}
	wg.Wait()
	finalVal, err := r.rdb.Get("cnt").Int()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Final value:", finalVal)
}

func singleMachineTest() {
	redisClient := NewClient()
	redisClient.syncProblem()
	redisClient.method1()
	redisClient.method2()
}
