package main

import (
	"context"
	"fmt"
)

type UserInfo struct {
	Name string
	Age  int
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "name", UserInfo{
		Name: "张三",
		Age:  18,
	})
	GetUser(ctx)
}

func GetUser(ctx context.Context) {
	fmt.Println(ctx.Value("name").(UserInfo).Name)
}
