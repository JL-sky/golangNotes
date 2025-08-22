package main

import (
	"encoding/json"
	"fmt"

	"github.com/r3labs/diff"
)

func compareJSON3(json1, json2 string) {
	var obj1, obj2 map[string]interface{}
	json.Unmarshal([]byte(json1), &obj1)
	json.Unmarshal([]byte(json2), &obj2)

	changelog, _ := diff.Diff(obj1, obj2)
	for _, change := range changelog {
		fmt.Printf("Path: %v\n", change.Path)
		fmt.Printf("From: %v\n", change.From)
		fmt.Printf("To: %v\n\n", change.To)
	}
}
