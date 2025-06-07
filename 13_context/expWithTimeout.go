package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 设置一个1s的超时上下文，用于控制所有goroutine的超时操作
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go HelloHandle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("Hello Handle ", ctx.Err())
	}
}

func HelloHandle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("Hello 123", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
