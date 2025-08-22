package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wI2L/jsondiff"
)

func compareJSON4(before, after string) {
	// 解析 JSON
	var beforeObj, afterObj interface{}
	if err := json.Unmarshal([]byte(before), &beforeObj); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal([]byte(after), &afterObj); err != nil {
		log.Fatal(err)
	}

	// 比较 JSON 差异
	patch, err := jsondiff.Compare(beforeObj, afterObj)
	if err != nil {
		log.Fatal(err)
	}

	// 处理差异结果
	fmt.Println("变更详情:")
	for _, op := range patch {
		fmt.Printf("- 操作类型: %s\n", op.Type)
		fmt.Printf("  路径: %s\n", op.Path)
		fmt.Printf("  原值: %v\n", op.OldValue)
		fmt.Printf("  新值: %v\n", op.Value)
		fmt.Println("------")
	}
}
