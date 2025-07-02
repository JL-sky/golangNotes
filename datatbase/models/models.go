package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Title       string
	Description string
	URL         string
}

type VideoHistory struct {
	gorm.Model
	VideoID     uint // 关联的原始视频ID
	Title       string
	Description string
	URL         string
	IsCurrent   bool // 标记是否为当前最新记录
}

type VideoChange struct {
	gorm.Model
	VideoID    uint      // 关联的原始视频ID
	ChangedAt  time.Time // 变更时间
	ChangeType string    // 变更类型: create/update/delete
	Field      string    // 变更的字段名
	OldValue   string    // 旧值
	NewValue   string    // 新值
}
