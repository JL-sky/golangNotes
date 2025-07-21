package changelog

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jl-sky/grom/golangNotes/datatbase/config"
	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// OutputImpl 变更处理器实现
type OutputImpl struct {
	db *gorm.DB
}

func NewOutput(db *gorm.DB) (*OutputImpl, error) {
	outputImpl := &OutputImpl{db: db}
	if err := outputImpl.InitOutput(); err != nil {
		return nil, fmt.Errorf("output init error: %w", err)
	}
	return outputImpl, nil
}

func (o *OutputImpl) InitOutput() error {
	tableName := GetTableNameWithHistory(config.TOutput)
	if HasTable(o.db, tableName) {
		return nil
	}

	// 创建历史表
	err := o.db.Table(tableName).AutoMigrate(&models.HistoryOutput{})
	if err != nil {
		return fmt.Errorf("failed to migrate history table: %w", err)
	}

	log.Infof("Created history table: %s", tableName)
	return nil
}

func (o *OutputImpl) Changelog(ctx context.Context, req *models.ChangelogAdminReq) error {
	if req == nil {
		return fmt.Errorf("changlog req is empty")
	}
	// 解析当前记录
	curRecord, err := o.getRowData(req.RowData)
	if err != nil {
		return fmt.Errorf("failed to parse row data: %w", err)
	}
	log.Debugf("curRecord is %+v", curRecord)
	// 查询历史记录
	preRecord, err := o.QueryHistoryData(ctx, curRecord)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("No history record found, handling first change")
			return o.handleFirstChange(ctx, curRecord)
		}
		return fmt.Errorf("failed to query history data: %w", err)
	}
	log.Debugf("preRecord is %+v", preRecord)
	return o.handleUpdate(ctx, preRecord, curRecord, req.ChangeUser)
}

func (o *OutputImpl) getRowData(rowData string) (*models.HistoryOutput, error) {
	var raw map[string]interface{}
	if err := json.Unmarshal([]byte(rowData), &raw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal row data: %w", err)
	}

	if config, ok := raw["c_config"].(map[string]interface{}); ok {
		configBytes, err := json.Marshal(config)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal c_config: %w", err)
		}
		raw["c_config"] = string(configBytes)
	}

	var result models.HistoryOutput
	if err := mapToStruct(raw, &result); err != nil {
		return nil, fmt.Errorf("failed to convert map to struct: %w", err)
	}

	return &result, nil
}

func mapToStruct(m map[string]interface{}, s interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, s)
}

func (o *OutputImpl) QueryHistoryData(ctx context.Context, curData *models.HistoryOutput) (*models.HistoryOutput, error) {
	if curData == nil {
		return nil, fmt.Errorf("current data cannot be nil")
	}
	hisTableName := GetTableNameWithHistory(config.TOutput)
	// 验证查询条件
	if curData.CFamilyID == "" {
		return nil, fmt.Errorf("familyID cannot be empty")
	}
	if curData.CDataType != "platform" {
		return nil, fmt.Errorf("invalid data type, expected 'platform' got '%s'", curData.CDataType)
	}
	if curData.CDataTypeValue == "" {
		return nil, fmt.Errorf("platform value cannot be empty")
	}

	// 查询当前有效记录
	var history models.HistoryOutput
	err := o.db.WithContext(ctx).
		Table(hisTableName).
		Where("c_family_id = ? AND c_data_type_value = ? AND record_status = ?",
							curData.CFamilyID, curData.CDataTypeValue, "current").
		Set("gorm:query_option", "FOR UPDATE"). // 添加行锁
		First(&history).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no current record found for familyID %s and platform %s: %w",
				curData.CFamilyID, curData.CDataTypeValue, err)
		}
		return nil, fmt.Errorf("failed to query history data: %w", err)
	}

	// 将HistoryOutput转换为TOutput
	return &history, nil
}

func (o *OutputImpl) handleFirstChange(ctx context.Context, curRecord *models.HistoryOutput) error {
	historyTable := GetTableNameWithHistory(config.TOutput)
	newHistory := models.HistoryOutput{
		CModifier:      curRecord.CModifier,
		CStrategyID:    curRecord.CStrategyID,
		CFamilyID:      curRecord.CFamilyID,
		CParentID:      curRecord.CParentID,
		CDataType:      curRecord.CDataType,
		CDataTypeValue: curRecord.CDataTypeValue,
		CStatus:        curRecord.CStatus,
		CMtime:         curRecord.CMtime,
		CPageUUID:      curRecord.CPageUUID,
		CCreator:       curRecord.CCreator,
		CVersion:       curRecord.CVersion,
		CVersionType:   curRecord.CVersionType,
		CPageClass:     curRecord.CPageClass,
		CRelationID:    curRecord.CRelationID,
		CIsPatch:       curRecord.CIsPatch,
		CType:          curRecord.CType,
		CConfig:        curRecord.CConfig,
		Changelogs:     "",
		RecordStatus:   "current",
	}
	if err := o.db.WithContext(ctx).
		Table(historyTable).
		Create(&newHistory).Error; err != nil {
		return fmt.Errorf("failed to create first history record: %w", err)
	}
	log.Infof("Created first history record for familyID %s and platform %s",
		curRecord.CFamilyID, curRecord.CDataTypeValue)
	return nil
}

func (o *OutputImpl) handleUpdate(ctx context.Context, preRecord, curRecord *models.HistoryOutput, changeUser string) error {
	log.Debugf("Previous record found, handling update")
	// 字段校验
	if preRecord.CFamilyID != curRecord.CFamilyID || preRecord.CDataTypeValue != curRecord.CDataTypeValue {
		log.Debug("familyId or dataTypeValue mismatch, skipping")
		return nil
	}
	// 比较差异
	change, err := CompareWithDiff(preRecord, curRecord)
	if err != nil {
		return fmt.Errorf("failed to compare records: %w", err)
	}
	if change == "" {
		log.Debug("No significant changes detected")
		return nil
	}
	log.Infof("Detected changes for familyID %s and platform %s: %+v",
		curRecord.CFamilyID, curRecord.CDataTypeValue, change)
	// 记录变更日志
	if err := o.RecordChangeLog(ctx, preRecord, curRecord, changeUser, change); err != nil {
		log.Errorf("Failed to record change log: %v", err)
	}

	//更新历史记录
	return o.db.Transaction(func(tx *gorm.DB) error {
		historyTable := GetTableNameWithHistory(config.TOutput)
		// 标记旧记录为历史版本
		if err := tx.WithContext(ctx).
			Table(historyTable).
			Where("c_family_id = ? AND c_data_type_value = ? AND record_status = ?",
				preRecord.CFamilyID, preRecord.CDataTypeValue, "current").
			Update("record_status", "historical").Error; err != nil {
			return fmt.Errorf("failed to mark old record as historical: %w", err)
		}
		// 创建新当前记录
		newHistory := models.HistoryOutput{
			CModifier:      changeUser,
			CStrategyID:    curRecord.CStrategyID,
			CFamilyID:      curRecord.CFamilyID,
			CParentID:      curRecord.CParentID,
			CDataType:      curRecord.CDataType,
			CDataTypeValue: curRecord.CDataTypeValue,
			CStatus:        curRecord.CStatus,
			CMtime:         time.Now().Format("2006-01-02 15:04:05"),
			CPageUUID:      curRecord.CPageUUID,
			CCreator:       curRecord.CCreator,
			CVersion:       curRecord.CVersion,
			CVersionType:   curRecord.CVersionType,
			CPageClass:     curRecord.CPageClass,
			CRelationID:    curRecord.CRelationID,
			CIsPatch:       curRecord.CIsPatch,
			CType:          curRecord.CType,
			CConfig:        curRecord.CConfig,
			RecordStatus:   "current",
		}
		if err := tx.WithContext(ctx).
			Table(historyTable).
			Create(&newHistory).Error; err != nil {
			return fmt.Errorf("failed to create new history record: %w", err)
		}
		log.Infof("Updated history record for familyID %s and platform %s",
			curRecord.CFamilyID, curRecord.CDataTypeValue)
		return nil
	})
}

// RecordChangeLog 记录变更日志并发送通知
func (o *OutputImpl) RecordChangeLog(ctx context.Context, preRecord, curRecord *models.HistoryOutput, changeUser, change string) error {
	if err := o.createChangeLog(ctx, curRecord, change); err != nil {
		return err
	}
	return nil
}

// createChangeLog 创建变更日志记录
func (o *OutputImpl) createChangeLog(ctx context.Context, record *models.HistoryOutput, change string) error {
	tableName := GetTableNameWithHistory(config.TOutput)

	changeLog := models.HistoryOutput{
		CStrategyID:    record.CStrategyID,
		CFamilyID:      record.CFamilyID,
		CParentID:      record.CParentID,
		CDataType:      record.CDataType,
		CDataTypeValue: record.CDataTypeValue,
		CStatus:        record.CStatus,
		CMtime:         time.Now().Format("2006-01-02 15:04:05"),
		CPageUUID:      record.CPageUUID,
		CCreator:       record.CCreator,
		CVersion:       record.CVersion,
		CVersionType:   record.CVersionType,
		CPageClass:     record.CPageClass,
		CRelationID:    record.CRelationID,
		CIsPatch:       record.CIsPatch,
		CType:          record.CType,
		CConfig:        record.CConfig,

		Changelogs:   change,
		RecordStatus: "current",
	}

	if err := o.db.WithContext(ctx).Table(tableName).Create(&changeLog).Error; err != nil {
		return fmt.Errorf("创建变更日志失败: %w", err)
	}
	return nil
}

func HasTable(db *gorm.DB, tableName string) bool {
	return db.Migrator().HasTable(tableName)
}

func GetTableNameWithHistory(tableName string) string {
	return tableName + "_history"
}
