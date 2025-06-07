package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wait = sync.WaitGroup{}

func main() {
	t1 := time.Now()
	ctx, cancel := context.WithCancel(context.Background())

	wait.Add(1)
	// 子协程获取ip地址
	go func() {
		ip, err := GetIp(ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ip)
		}
		wait.Done()
	}()

	// 现在需要另外一个协程在2s后取消所有操作
	go func() {
		time.Sleep(2 * time.Second)
		cancel() // 取消所有操作
	}()
	wait.Wait()
	fmt.Println("main", time.Since(t1))
}

func GetIp(ctx context.Context) (string, error) {
	// 创建通道用于接收结果
	resultCh := make(chan string)

	// 启动子协程执行耗时操作
	go func() {
		time.Sleep(4 * time.Second)
		resultCh <- "192.168.1.1"
	}()

	// 主流程通过select多路复用
	select {
	case <-ctx.Done():
		return "", ctx.Err() // 2秒时触发此分支
	case ip := <-resultCh:
		return ip, nil // 4秒时触发此分支（若未提前取消）
	}
}
