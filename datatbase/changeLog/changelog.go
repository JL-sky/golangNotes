package changelog

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jl-sky/grom/golangNotes/datatbase/config"
	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	"github.com/r3labs/diff"
	"gorm.io/gorm"
)

// 初始化变更记录系统
func InitChangeLogSystem(db *gorm.DB) error {
	if HasTable(db, GetTableNameWithHistory(config.TChangeLogs)) {
		return nil
	}
	err := db.AutoMigrate(&models.TChangeLogs{})
	if err != nil {
		return fmt.Errorf("output init error")
	}
	return nil
}

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

		// 跳过时间类型的比较
		fromType := reflect.TypeOf(change.From)
		if fromType == reflect.TypeOf(time.Time{}) {
			continue
		}

		// 处理变更前的值
		if change.From != nil {
			beforeChanges[fieldName] = change.From
		} else {
			beforeChanges[fieldName] = nil
		}

		// 处理变更后的值
		if change.To != nil {
			afterChanges[fieldName] = change.To
		} else {
			afterChanges[fieldName] = nil
		}
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

		return string(jsonData), nil
	}

	return "", nil
}
