package main

import (
	"fmt"
	"log"

	pb "github.com/jl-sky/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func tets1() {
	// 创建不同类型的消息
	person := &pb.Person{
		Name: "Alice",
		Age:  30,
	}

	book := &pb.Book{
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
	}

	// 将消息转换为 anypb.Any
	personAny, err := anypb.New(person)
	if err != nil {
		log.Fatalf("Failed to create Any for Person: %v", err)
	}
	log.Printf("PersonAny: %v", personAny) // 输出: Person: {Name:Alice Age:30}

	bookAny, err := anypb.New(book)
	if err != nil {
		log.Fatalf("Failed to create Any for Book: %v", err)
	}

	// 创建 Container 并存储不同类型的消息
	container1 := &pb.Container{Content: personAny}
	container2 := &pb.Container{Content: bookAny}

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
	var parsedContainer pb.Container

	// 解析包含 Person 的容器
	if err := proto.Unmarshal(container1Data, &parsedContainer); err != nil {
		log.Fatalf("Failed to unmarshal container1: %v", err)
	}

	// 检查 content 类型并还原
	if parsedContainer.Content.MessageIs(&pb.Person{}) {
		var p pb.Person
		if err := parsedContainer.Content.UnmarshalTo(&p); err != nil {
			log.Fatalf("Failed to unmarshal Person: %v", err)
		}
		fmt.Printf("Got Person: %+v\n", p) // 输出: Got Person: {Name:Alice Age:30}
	}

	// 解析包含 Book 的容器
	if err := proto.Unmarshal(container2Data, &parsedContainer); err != nil {
		log.Fatalf("Failed to unmarshal container2: %v", err)
	}

	if parsedContainer.Content.MessageIs(&pb.Book{}) {
		var b pb.Book
		if err := parsedContainer.Content.UnmarshalTo(&b); err != nil {
			log.Fatalf("Failed to unmarshal Book: %v", err)
		}
		fmt.Printf("Got Book: %+v\n", b) // 输出: Got Book: {Title:The Go Programming Language Author:Alan Donovan}
	}
}

func tets2() {
	// 创建 Person 消息
	person := &pb.Person{
		Name: "Alice",
		Age:  30,
	}
	// 将消息转换为 anypb.Any
	personAny, err := anypb.New(person)
	if err != nil {
		log.Fatalf("Failed to create Any for Person: %v", err)
	}
	// 创建 Container 并存储 Person 消息
	container := &pb.Container{Content: personAny}

	var p pb.Person
	// 解析 anypb.Any 并还原为 Person 消息
	if err := container.Content.UnmarshalTo(&p); err != nil {
		log.Fatalf("Failed to unmarshal Person: %v", err)
	}
	fmt.Printf("Method 3: Name=%q, Age=%d\n", p.GetName(), p.GetAge())

}

func main() {
	tets2()
}
