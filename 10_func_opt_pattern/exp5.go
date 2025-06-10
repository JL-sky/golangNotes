package main

import (
	"fmt"
	"log"
	"strings"
)

type Request struct {
	query string
	data  string
}

func BuildRequestByTraceId(traceId string) *Request {
	q := fmt.Sprintf("traceid=%s", traceId)
	request := &Request{
		query: q,
		data:  "",
	}
	log.Printf("request: %v", request)
	return request
}

func BuildRequest(qimei string) *Request {
	q := fmt.Sprintf("qimei=%s", qimei)
	request := &Request{
		query: q,
		data:  "",
	}
	log.Printf("request: %v", request)
	return request
}

type queryOption func(*Request)

func WithQimei(qimei string) queryOption {
	return func(r *Request) {
		if qimei != "" {
			r.query = fmt.Sprintf("qimei=%s", qimei)
		}
	}
}

func WithTraceId(traceId string) queryOption {
	return func(r *Request) {
		if traceId != "" {
			r.query = fmt.Sprintf("traceid=%s", traceId)
		}
	}
}

// WithData 设置请求数据的选项函数
func WithData(data string) queryOption {
	return func(r *Request) {
		r.data = data
	}
}

func NewRequest(options ...queryOption) *Request {
	request := &Request{}
	for _, opt := range options {
		opt(request)
	}
	log.Printf("request: %v", request)
	return request
}

func GetPage(request *Request) string {
	if request == nil {
		return ""
	}
	if request.query == "" {
		return ""
	}
	if strings.Contains(request.query, "qimei") {
		return "qimei"
	}
	if strings.Contains(request.query, "traceid") {
		return "traceid"
	}

	return ""
}

func Query(url string) string {
	// 使用函数选项模式构建请求
	request := NewRequest(WithQimei("123456789"), WithData("some data"))
	return GetPage(request)
}

func main() {
	// 示例：构建不同类型的请求
	req1 := NewRequest(WithQimei("abcdefg"))
	req2 := NewRequest(WithTraceId("12345"), WithData("test"))

	log.Printf("req1: %v", req1)
	log.Printf("req2: %v", req2)
}
