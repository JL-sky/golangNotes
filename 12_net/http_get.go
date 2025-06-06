package main

import (
	"fmt"
	"io"
	"net/http"
)

func test1(url string) {
	// 发送http get请求
	rsp, err := http.Get(url)
	// 如果发生错误，打印错误信息并返回
	if err != nil {
		fmt.Println("http get err:", err)
		return
	}
	// 关闭http响应体
	defer rsp.Body.Close()
	// 读取http响应体内容
	content, err := io.ReadAll(rsp.Body)
	// 打印http响应体内容
	fmt.Println(string(content))
}

func test2(url string) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", "123456")
	resp, _ := client.Do(request)
	content, _ := io.ReadAll(resp.Body)
	fmt.Println(string(content))
	defer resp.Body.Close()
}

func main() {
	url := "https://www.baidu.com"
	// test1(url)
	test2(url)
}
