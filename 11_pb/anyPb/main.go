package main

import (
	"fmt"
	"log"

	example "github.com/jl-sky/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	// 创建不同类型的消息
	person := &example.Person{
		Name: "Alice",
		Age:  30,
	}

	book := &example.Book{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
	}

	// 将消息转换为 anypb.Any
	personAny, err := anypb.New(person)
	if err != nil {
		log.Fatalf("Failed to create Any for Person: %v", err)
	}

	bookAny, err := anypb.New(book)
	if err != nil {
		log.Fatalf("Failed to create Any for Book: %v", err)
	}

	// 创建 Container 并存储不同类型的消息
	container1 := &example.Container{Content: personAny}
	container2 := &example.Container{Content: bookAny}

	// 序列化 Container 消息
	container1Data, err := proto.Marshal(container1)
	if err != nil {
		log.Fatalf("Failed to marshal container1: %v", err)
	}

	container2Data, err := proto.Marshal(container2)
	if err != nil {
		log.Fatalf("Failed to marshal container2: %v", err)
	}

	// 反序列化并还原原始消息类型
	var parsedContainer example.Container

	// 解析包含 Person 的容器
	if err := proto.Unmarshal(container1Data, &parsedContainer); err != nil {
		log.Fatalf("Failed to unmarshal container1: %v", err)
	}

	// 检查 content 类型并还原
	if parsedContainer.Content.MessageIs(&example.Person{}) {
		var p example.Person
		if err := parsedContainer.Content.UnmarshalTo(&p); err != nil {
			log.Fatalf("Failed to unmarshal Person: %v", err)
		}
		fmt.Printf("Got Person: %+v\n", p) // 输出: Got Person: {Name:Alice Age:30}
	}

	// 解析包含 Book 的容器
	if err := proto.Unmarshal(container2Data, &parsedContainer); err != nil {
		log.Fatalf("Failed to unmarshal container2: %v", err)
	}

	if parsedContainer.Content.MessageIs(&example.Book{}) {
		var b example.Book
		if err := parsedContainer.Content.UnmarshalTo(&b); err != nil {
			log.Fatalf("Failed to unmarshal Book: %v", err)
		}
		fmt.Printf("Got Book: %+v\n", b) // 输出: Got Book: {Title:The Go Programming Language Author:Alan Donovan}
	}
}
