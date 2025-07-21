package changelog

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jl-sky/grom/golangNotes/datatbase/config"
	"github.com/r3labs/diff"
	log "github.com/sirupsen/logrus"
)

// CompareWithDiff 比较两个结构体的差异
func CompareWithDiff(a, b interface{}) (string, error) {
	changelog, err := diff.Diff(a, b)
	if err != nil {
		return "", fmt.Errorf("diff comparison failed: %v", err)
	}

	// 创建变更前后的map
	beforeChanges := make(map[string]interface{})
	afterChanges := make(map[string]interface{})

	for _, change := range changelog {
		if len(change.Path) == 0 {
			continue
		}

		// 获取字段名，处理嵌套路径如"User.Name"
		fieldName := strings.Join(change.Path, ".")
		if _, ok := config.FilterFields[fieldName]; ok {
			log.Debugf("过滤字段: %s", fieldName)
			continue
		}
		// 跳过时间类型的比较
		fromType := reflect.TypeOf(change.From)
		if fromType == reflect.TypeOf(time.Time{}) {
			continue
		}

		// 处理变更前的值
		log.Debugf("变更字段: %s", fieldName)
		beforeChanges[fieldName] = change.From

		// 处理变更后的值
		afterChanges[fieldName] = change.To
	}

	// 如果有变更，生成最终结果
	if len(beforeChanges) > 0 || len(afterChanges) > 0 {
		result := map[string]interface{}{
			"before": beforeChanges,
			"after":  afterChanges,
		}

		jsonData, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return "", fmt.Errorf("JSON序列化失败: %v", err)
		}
		log.Debugf(string(jsonData))
		return string(jsonData), nil
	}

	return "", nil
}
