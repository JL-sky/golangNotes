package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 处理GET请求
func getHandler(w http.ResponseWriter, r *http.Request) {
	// 检查请求方法是否为GET
	if r.Method == http.MethodGet {
		// 获取URL查询参数
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		fmt.Fprintf(w, "Hello, %s! This is a GET request.", name)
	}
}

// 处理POST请求
func postHandler(w http.ResponseWriter, r *http.Request) {
	// 检查请求方法是否为POST
	if r.Method == http.MethodPost {
		// 读取请求体内容
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusInternalServerError)
			return
		}
		// 返回响应
		fmt.Fprintf(w, "Received POST request with body: %s", string(body))
	}
}

func main() {
	// 设置路由
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)

	// 启动HTTP服务器
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
