package main

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	log "github.com/sirupsen/logrus"
)

func changelogAdminTest() {
	req := mockReq()
	ChangeLogByAdmin(context.Background(), req)
}
func mockReq() *models.ChangelogAdminReq {
	// rowData := "{\"c_modifier\":\"cheersjiang\",\"c_strategy_id\":\"9d115c1bad99493db72245b27df89885\",\"c_family_id\":\"3019164517a844e89e7f23699744f3ee\",\"c_parent_id\":\"3019164517a844e89e7f23699744f3ee\",\"c_data_type\":\"platform\",\"c_data_type_value\":\"harmony_phone\",\"c_status\":1,\"c_mtime\":\"2025-07-10 20:27:58\",\"c_page_uuid\":\"cc54892bd8964078900eeff6b1bdc969\",\"c_creator\":\"sennazang\",\"c_version\":\"81\",\"c_version_type\":\"release\",\"c_page_class\":\"\",\"c_relation_id\":\"\",\"c_is_patch\":false,\"c_type\":\"default\",\"c_config\":\"config\"}"
	changeLogAdmin := &models.ChangelogAdminReq{
		TableName:  "t_output",
		ChangeUser: "mvl_debug",
		RowData:    rowData,
	}
	return changeLogAdmin
}

// TestConcurrentUpdates 测试不同平台的并发变更
func TestConcurrentUpdates() {
	// 准备基础测试数据
	baseData := map[string]interface{}{
		"c_family_id":    "family_123", // 固定familyID
		"c_data_type":    "platform",
		"c_status":       1,
		"c_mtime":        time.Now().Format("2006-01-02 15:04:05"),
		"c_page_uuid":    "page_123",
		"c_creator":      "test_user",
		"c_version":      "1.0",
		"c_version_type": "release",
		"c_config":       map[string]interface{}{"key": "value"},
	}

	// 定义要测试的不同平台
	platforms := []string{"mac", "pc", "harmony", "android"}

	// 准备并发测试的请求
	var requests []*models.ChangelogAdminReq
	for _, platform := range platforms {
		// 复制基础数据
		data := make(map[string]interface{})
		for k, v := range baseData {
			data[k] = v
		}
		// 设置不同的dataTypeValue
		data["c_data_type_value"] = platform

		// 转换为JSON字符串
		rowData, err := json.Marshal(data)
		if err != nil {
			log.Errorf("Failed to marshal data for platform %s: %v", platform, err)
			continue
		}

		requests = append(requests, &models.ChangelogAdminReq{
			TableName:  "t_output",
			ChangeUser: "tester",
			RowData:    string(rowData),
		})
	}

	// 模拟并发请求
	var wg sync.WaitGroup
	for i, req := range requests {
		wg.Add(1)
		go func(id int, request *models.ChangelogAdminReq) {
			defer wg.Done()

			ctx := context.Background()
			err := ChangeLogByAdmin(ctx, request)
			if err != nil {
				log.Errorf("Goroutine %d (%s) failed: %v",
					id, request.RowData, err)
			} else {
				log.Infof("Goroutine %d (%s) succeeded",
					id, request.RowData)
			}
		}(i, req)
	}

	wg.Wait()
	log.Info("All concurrent updates completed")
}
