package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

func syncDisTest() {
	// client := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })

	// // 测试无锁情况
	// testWithoutLock(client)

	// // 测试RedLock解决方案
	// testWithRedLock(client)

	// 测试多节点RedLock解决方案
	testMultiNodeWithRedLock()
}

func testWithoutLock(client *redis.Client) {
	ctx := context.Background()
	key := "counter"
	client.Set(ctx, key, 0, 0)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val, _ := client.Get(ctx, key).Int()
			val++
			client.Set(ctx, key, val, 0)
		}()
	}
	wg.Wait()

	finalVal, _ := client.Get(ctx, key).Int()
	fmt.Printf("无锁情况下最终计数: %d (应得100)\n", finalVal)
}

// 每个协程模拟一个机器节点，使用RedLock解决分布式锁问题
func testWithRedLock(client *redis.Client) {
	ctx := context.Background()
	key := "counter"
	client.Set(ctx, key, 0, 0)

	// 创建RedLock实例
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 创建Mutex实例
			mutex := rs.NewMutex("counter-lock",
				// 设置Mutex过期时间为10秒
				redsync.WithExpiry(10*time.Second),
				// 设置Mutex尝试次数为10
				redsync.WithTries(10),
				// 设置Mutex重试间隔为100毫秒
				redsync.WithRetryDelay(100*time.Millisecond),
			)

			// 带重试的锁获取
			for retry := 0; retry < 3; retry++ {
				// 尝试获取锁
				if err := mutex.Lock(); err != nil {
					fmt.Printf("Goroutine %d failed to acquire lock (attempt %d): %v\n",
						id, retry+1, err)
					/*
						​指数退避等待​​：time.Sleep(time.Duration(retry+1) * 100 * time.Millisecond)
							第一次重试等待 100ms (retry=0→ 0+1=1 → 1×100ms)
							第二次重试等待 200ms (retry=1→ 1+1=2 → 2×100ms)
							第三次重试等待 300ms (retry=2→ 2+1=3 → 3×100ms)（这种递增等待时间的设计称为 ​​指数退避​​，有助于减少竞争）
					*/
					time.Sleep(time.Duration(retry+1) * 100 * time.Millisecond)
					continue
				}

				// 临界区
				val, _ := client.Get(ctx, key).Int()
				val++
				client.Set(ctx, key, val, 0)

				mutex.Unlock()
				return
			}
		}(i)
	}
	wg.Wait()

	finalVal, _ := client.Get(ctx, key).Int()
	fmt.Printf("使用RedLock后最终计数: %d\n", finalVal)
}

func testMultiNodeWithRedLock() {
	// 1. 初始化多个Redis节点
	nodes := []*redis.Client{
		redis.NewClient(&redis.Options{Addr: "localhost:6379"}),
		redis.NewClient(&redis.Options{Addr: "localhost:6380"}),
		redis.NewClient(&redis.Options{Addr: "localhost:6381"}),
	}

	// 2. 创建RedSync实例
	pools := make([]redsyncredis.Pool, len(nodes))
	for i, node := range nodes {
		pools[i] = goredis.NewPool(node)
	}
	rs := redsync.New(pools...)

	// 3. 在所有节点初始化计数器（使用主节点）
	ctx := context.Background()
	key := "global_counter"
	nodes[0].Set(ctx, key, 0, 0) // 只需要在一个节点初始化

	// 4. 模拟多节点多worker
	var wg sync.WaitGroup
	workersPerNode := 10

	for nodeID := 0; nodeID < len(nodes); nodeID++ {
		for workerID := 0; workerID < workersPerNode; workerID++ {
			wg.Add(1)
			go func(nodeID, workerID int) {
				defer wg.Done()

				// 获取分布式锁（保护全局资源）
				mutex := rs.NewMutex("global_lock",
					redsync.WithExpiry(2*time.Second),
					redsync.WithTries(10),
				)

				if err := mutex.Lock(); err != nil {
					fmt.Printf("Node %d Worker %d lock failed: %v\n", nodeID, workerID, err)
					return
				}
				defer mutex.Unlock()

				// 所有worker都操作同一个主节点（或使用WATCH/MULTI实现跨节点原子操作）
				client := nodes[0]

				// 使用INCR保证原子性
				if err := client.Incr(ctx, key).Err(); err != nil {
					fmt.Printf("Node %d Worker %d incr failed: %v\n", nodeID, workerID, err)
					return
				}

				fmt.Printf("Node %d Worker %d incr success\n", nodeID, workerID)
			}(nodeID, workerID)
		}
	}

	wg.Wait()

	// 5. 验证结果（从所有节点读取，应该相同）
	for i, node := range nodes {
		val, _ := node.Get(ctx, key).Int()
		fmt.Printf("Node %d final count: %d (expected %d)\n",
			i, val, workersPerNode*len(nodes))
	}
}
