package changelog

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	"github.com/jl-sky/grom/golangNotes/datatbase/mysql"
	"gorm.io/gorm"
)

// 初始化变更记录系统
func InitChangeLogSystem(db *gorm.DB) error {
	// 自动迁移所有相关表
	err := db.AutoMigrate(
		&models.Video{},
		&models.VideoHistory{},
		&models.VideoChange{},
	)
	if err != nil {
		return fmt.Errorf("自动迁移表失败: %v", err)
	}
	return nil
}

// 处理变更记录的接口
func HandleVideoChange(db *gorm.DB, currentVideo *models.Video) error {
	// 0. 确保系统已初始化
	if err := checkAndInitTables(db); err != nil {
		return err
	}

	// 1. 检查是否为新增记录
	if currentVideo.ID == 0 {
		return handleNewVideo(db, currentVideo)
	}

	// 2. 获取上一条历史记录
	var previousRecord models.VideoHistory
	err := db.Where("video_id = ? AND is_current = true", currentVideo.ID).First(&previousRecord).Error

	// 3. 处理不同情况
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 没有历史记录，这是第一次变更
		return handleFirstChange(db, currentVideo)
	} else if err != nil {
		return fmt.Errorf("查询历史记录失败: %v", err)
	}

	// 4. 对比变更并记录差异
	return handleUpdate(db, currentVideo, &previousRecord)
}

// 检查并初始化表
func checkAndInitTables(db *gorm.DB) error {
	// 检查表是否存在，不存在则创建
	if !db.Migrator().HasTable(&models.Video{}) {
		if err := db.Migrator().CreateTable(&models.Video{}); err != nil {
			return fmt.Errorf("创建Video表失败: %v", err)
		}
	}

	if !db.Migrator().HasTable(&models.VideoHistory{}) {
		if err := db.Migrator().CreateTable(&models.VideoHistory{}); err != nil {
			return fmt.Errorf("创建VideoHistory表失败: %v", err)
		}
	}

	if !db.Migrator().HasTable(&models.VideoChange{}) {
		if err := db.Migrator().CreateTable(&models.VideoChange{}); err != nil {
			return fmt.Errorf("创建VideoChange表失败: %v", err)
		}
	}

	return nil
}

// 处理新增视频
func handleNewVideo(db *gorm.DB, video *models.Video) error {
	// 保存原始记录
	if err := mysql.CreateVideo(db, video); err != nil {
		return fmt.Errorf("创建视频失败: %v", err)
	}

	// 创建历史记录
	history := models.VideoHistory{
		VideoID:     video.ID,
		Title:       video.Title,
		Description: video.Description,
		URL:         video.URL,
		IsCurrent:   true,
	}

	if err := db.Create(&history).Error; err != nil {
		return fmt.Errorf("创建历史记录失败: %v", err)
	}

	// 记录创建变更
	change := models.VideoChange{
		VideoID:    video.ID,
		ChangedAt:  time.Now(),
		ChangeType: "create",
		Field:      "all",
		NewValue:   fmt.Sprintf("视频创建: %s", video.Title),
	}

	return db.Create(&change).Error
}

// 处理第一次变更
func handleFirstChange(db *gorm.DB, video *models.Video) error {
	// 创建历史记录
	history := models.VideoHistory{
		VideoID:     video.ID,
		Title:       video.Title,
		Description: video.Description,
		URL:         video.URL,
		IsCurrent:   true,
	}

	if err := db.Create(&history).Error; err != nil {
		return fmt.Errorf("创建历史记录失败: %v", err)
	}

	// 记录初始变更
	change := models.VideoChange{
		VideoID:    video.ID,
		ChangedAt:  time.Now(),
		ChangeType: "initial",
		Field:      "all",
		NewValue:   fmt.Sprintf("初始记录: %s", video.Title),
	}

	return db.Create(&change).Error
}

// 处理更新变更
func handleUpdate(db *gorm.DB, currentVideo *models.Video, previousRecord *models.VideoHistory) error {
	// 1. 标记旧历史记录为非当前
	previousRecord.IsCurrent = false
	if err := db.Save(previousRecord).Error; err != nil {
		return fmt.Errorf("更新历史记录状态失败: %v", err)
	}

	// 2. 创建新历史记录
	newHistory := models.VideoHistory{
		VideoID:     currentVideo.ID,
		Title:       currentVideo.Title,
		Description: currentVideo.Description,
		URL:         currentVideo.URL,
		IsCurrent:   true,
	}
	if err := db.Create(&newHistory).Error; err != nil {
		return fmt.Errorf("创建新历史记录失败: %v", err)
	}

	// 3. 对比字段变化并记录
	var changes []models.VideoChange
	// changes := compareAndRecordChanges(currentVideo, previousRecord)
	if previousRecord.Title != currentVideo.Title {
		changes = append(changes, models.VideoChange{
			VideoID:    currentVideo.ID,
			ChangedAt:  time.Now(),
			ChangeType: "update",
			Field:      "title",
			OldValue:   previousRecord.Title,
			NewValue:   currentVideo.Title,
		})
	}

	if previousRecord.Description != currentVideo.Description {
		changes = append(changes, models.VideoChange{
			VideoID:    currentVideo.ID,
			ChangedAt:  time.Now(),
			ChangeType: "update",
			Field:      "description",
			OldValue:   previousRecord.Description,
			NewValue:   currentVideo.Description,
		})
	}

	if previousRecord.URL != currentVideo.URL {
		changes = append(changes, models.VideoChange{
			VideoID:    currentVideo.ID,
			ChangedAt:  time.Now(),
			ChangeType: "update",
			Field:      "url",
			OldValue:   previousRecord.URL,
			NewValue:   currentVideo.URL,
		})
	}

	// 4. 保存所有变更记录
	if len(changes) > 0 {
		return db.Create(&changes).Error
	}

	return nil
}

// 比较并记录变更
func compareAndRecordChanges(currentVideo *models.Video, previousRecord *models.VideoHistory) []models.VideoChange {
	var changes []models.VideoChange
	now := time.Now()

	currentValue := reflect.ValueOf(currentVideo).Elem()
	previousValue := reflect.ValueOf(previousRecord).Elem()
	currentType := currentValue.Type()

	for i := 0; i < currentValue.NumField(); i++ {
		field := currentType.Field(i)
		fieldName := field.Name

		// 跳过不需要比较的字段
		if fieldName == "Model" || fieldName == "ID" || fieldName == "CreatedAt" ||
			fieldName == "UpdatedAt" || fieldName == "DeletedAt" || fieldName == "IsCurrent" {
			continue
		}

		currentField := currentValue.Field(i)
		previousField := previousValue.Field(i)

		// 比较字段值
		if !reflect.DeepEqual(currentField.Interface(), previousField.Interface()) {
			changes = append(changes, models.VideoChange{
				VideoID:    currentVideo.ID,
				ChangedAt:  now,
				ChangeType: "update",
				Field:      strings.ToLower(fieldName),
				OldValue:   fmt.Sprintf("%v", previousField.Interface()),
				NewValue:   fmt.Sprintf("%v", currentField.Interface()),
			})
		}
	}

	return changes
}
