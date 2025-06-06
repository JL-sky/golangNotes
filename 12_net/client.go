package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPage(url string) {
	method := "get"
	name := "GoUser"
	url = fmt.Sprintf("%s/%s?name=%s", url, method, name)
	fmt.Println(url)
	// 发送GET请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印响应内容
	fmt.Println("Response from GET:", string(body))
}

func PostPage(url string) {
	url = fmt.Sprintf("%s/post", url)
	// 创建POST请求的内容
	data := []byte(`{"message": "Hello, Server!"}`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印响应内容
	fmt.Println("Response from POST:", string(body))
}

func main() {
	url := "http://localhost:8080"
	GetPage(url)
	PostPage(url)
}
