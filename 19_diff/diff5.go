package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/wI2L/jsondiff"
)

type Config struct {
	CConfig  string `json:"CConfig"`
	CVersion string `json:"CVersion"`
}

func compare5(before, after string) {
	// 解析顶级 JSON
	var beforeConfig, afterConfig Config
	if err := json.Unmarshal([]byte(before), &beforeConfig); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal([]byte(after), &afterConfig); err != nil {
		log.Fatal(err)
	}

	// 解析嵌套的 CConfig JSON
	var beforeCConfig, afterCConfig interface{}
	if err := json.Unmarshal([]byte(beforeConfig.CConfig), &beforeCConfig); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal([]byte(afterConfig.CConfig), &afterCConfig); err != nil {
		log.Fatal(err)
	}

	// 比较顶级字段
	fmt.Println("比较顶级字段差异:")
	compareAndPrint(beforeConfig, afterConfig)

	// 比较 CConfig 内部字段
	fmt.Println("\n比较 CConfig 内部字段差异:")
	compareAndPrint(beforeCConfig, afterCConfig)
}

func compareAndPrint(before, after interface{}) {
	patch, err := jsondiff.Compare(before, after)
	if err != nil {
		log.Fatal(err)
	}

	for _, op := range patch {
		fmt.Printf("- 操作类型: %s\n", op.Type)
		fmt.Printf("  路径: %s\n", op.Path)
		if op.OldValue != nil {
			fmt.Printf("  原值: %v\n", op.OldValue)
		}
		if op.Value != nil {
			fmt.Printf("  新值: %v\n", op.Value)
		}
		fmt.Println("------")
	}
}
