package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// DiffResult 存储差异结果
type DiffResult struct {
	Path    string      `json:"path"`    // 字段路径
	Value1  interface{} `json:"value1"`  // 第一个JSON中的值
	Value2  interface{} `json:"value2"`  // 第二个JSON中的值
	Message string      `json:"message"` // 差异描述
}

// CompareJSON 比较两个JSON字符串
func CompareJSON2(jsonStr1, jsonStr2 string) ([]DiffResult, error) {
	var obj1, obj2 interface{}

	// 解析JSON字符串
	if err := json.Unmarshal([]byte(jsonStr1), &obj1); err != nil {
		return nil, fmt.Errorf("解析第一个JSON失败: %v", err)
	}
	if err := json.Unmarshal([]byte(jsonStr2), &obj2); err != nil {
		return nil, fmt.Errorf("解析第二个JSON失败: %v", err)
	}

	// 比较两个对象
	var diffs []DiffResult
	compareObjects("", obj1, obj2, &diffs)
	return diffs, nil
}

// compareObjects 递归比较两个对象
func compareObjects(path string, a, b interface{}, diffs *[]DiffResult) {
	if reflect.DeepEqual(a, b) {
		return
	}

	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)

	// 类型不同
	if aValue.Kind() != bValue.Kind() {
		*diffs = append(*diffs, DiffResult{
			Path:    path,
			Value1:  a,
			Value2:  b,
			Message: fmt.Sprintf("类型不同: %T vs %T", a, b),
		})
		return
	}

	switch aValue.Kind() {
	case reflect.Map:
		aMap, ok := a.(map[string]interface{})
		if !ok {
			*diffs = append(*diffs, DiffResult{
				Path:    path,
				Value1:  a,
				Value2:  b,
				Message: "无法解析为map[string]interface{}",
			})
			return
		}
		bMap, ok := b.(map[string]interface{})
		if !ok {
			*diffs = append(*diffs, DiffResult{
				Path:    path,
				Value1:  a,
				Value2:  b,
				Message: "无法解析为map[string]interface{}",
			})
			return
		}

		// 检查a中有而b中没有的键
		for key, aVal := range aMap {
			newPath := buildPath(path, key)
			if bVal, exists := bMap[key]; !exists {
				*diffs = append(*diffs, DiffResult{
					Path:    newPath,
					Value1:  aVal,
					Value2:  nil,
					Message: "第二个JSON中缺少此字段",
				})
			} else {
				compareObjects(newPath, aVal, bVal, diffs)
			}
		}

		// 检查b中有而a中没有的键
		for key, bVal := range bMap {
			if _, exists := aMap[key]; !exists {
				newPath := buildPath(path, key)
				*diffs = append(*diffs, DiffResult{
					Path:    newPath,
					Value1:  nil,
					Value2:  bVal,
					Message: "第一个JSON中缺少此字段",
				})
			}
		}

	case reflect.Slice, reflect.Array:
		aSlice, ok := a.([]interface{})
		if !ok {
			*diffs = append(*diffs, DiffResult{
				Path:    path,
				Value1:  a,
				Value2:  b,
				Message: "无法解析为[]interface{}",
			})
			return
		}
		bSlice, ok := b.([]interface{})
		if !ok {
			*diffs = append(*diffs, DiffResult{
				Path:    path,
				Value1:  a,
				Value2:  b,
				Message: "无法解析为[]interface{}",
			})
			return
		}

		// 数组长度不同
		if len(aSlice) != len(bSlice) {
			*diffs = append(*diffs, DiffResult{
				Path:    path,
				Value1:  len(aSlice),
				Value2:  len(bSlice),
				Message: fmt.Sprintf("数组长度不同: %d vs %d", len(aSlice), len(bSlice)),
			})
		}

		// 比较数组元素
		minLen := len(aSlice)
		if len(bSlice) < minLen {
			minLen = len(bSlice)
		}
		for i := 0; i < minLen; i++ {
			newPath := fmt.Sprintf("%s[%d]", path, i)
			compareObjects(newPath, aSlice[i], bSlice[i], diffs)
		}

	default:
		// 基本类型直接比较值
		*diffs = append(*diffs, DiffResult{
			Path:    path,
			Value1:  a,
			Value2:  b,
			Message: "值不同",
		})
	}
}

// buildPath 构建字段路径
func buildPath(base, key string) string {
	if base == "" {
		return key
	}
	return base + "." + key
}

// formatDiffResult 格式化差异结果
func formatDiffResult(diffs []DiffResult) string {
	var sb strings.Builder
	for _, diff := range diffs {
		sb.WriteString(fmt.Sprintf("路径: %s\n", diff.Path))
		sb.WriteString(fmt.Sprintf("  第一个值: %v\n", diff.Value1))
		sb.WriteString(fmt.Sprintf("  第二个值: %v\n", diff.Value2))
		sb.WriteString(fmt.Sprintf("  差异描述: %s\n\n", diff.Message))
	}
	return sb.String()
}
