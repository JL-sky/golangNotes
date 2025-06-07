package main

import (
	"fmt"
	"net/http"
	"time"
)

func test() {
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":8080", nil)
}

func SayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(&request)

	// 监控goroutine
	go func() {
		for range time.Tick(time.Second) {
			select {
			case <-request.Context().Done():
				fmt.Println("request is outgoing")
				return
			default:
				fmt.Println("Current request is in progress")
			}
		}
	}()

	//模拟请求耗时
	time.Sleep(2 * time.Second)
	writer.Write([]byte("Hi"))
}

func main() {
	test()
}
