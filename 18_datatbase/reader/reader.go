package reader

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jl-sky/grom/golangNotes/datatbase/config"
	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	"github.com/jl-sky/grom/golangNotes/datatbase/mysql"
	"gorm.io/gorm"
)

// ChangLogHeaderInfo 变更日志头信息
type ChangLogHeaderInfo struct {
	TableName     string `json:"table_name"`      // 表名
	TableKey      string `json:"table_key"`       // 表键（格式为"key1:key2:key3..."）
	TableKeyValue string `json:"table_key_value"` // 表键值（格式为"value1:value2:value3..."）
	ChangeUser    string `json:"change_user"`     // 变更用户
	CheckTime     int64  `json:"check_time"`      // 变更时间戳
}

// NotificationResponse 变更通知响应
type NotificationResponse struct {
	Success bool   `json:"success"` // 是否成功
	Message string `json:"message"` // 返回消息
}

// ReaderImpl 实现变更日志处理接口
type ReaderImpl struct {
}

func NewReaderImpl() *ReaderImpl {
	return &ReaderImpl{}
}

// NotificationChan 全局通知通道
var NotificationChan = make(chan ChangLogHeaderInfo, 100)

// HandleNotification 处理变更通知
func (r *ReaderImpl) HandleNotification(ctx context.Context, header *ChangLogHeaderInfo) *NotificationResponse {
	// 1. 基础验证
	if header.TableName == "" {
		return &NotificationResponse{
			Success: false,
			Message: "table name is empty",
		}
	}

	if header.TableKey == "" {
		return &NotificationResponse{
			Success: false,
			Message: "table key is empty",
		}
	}

	// 2. 将通知发送到通道
	go func() {
		select {
		case NotificationChan <- *header:
			// 通知成功入队
		case <-time.After(1 * time.Second):
			// 超时处理，可以记录日志
			log.Printf("Notification queue timeout for table: %s", header.TableName)
		case <-ctx.Done():
			// 上下文取消
			return
		}
	}()

	// 3. 立即返回成功响应
	return &NotificationResponse{
		Success: true,
		Message: "notification received and queued successfully",
	}
}

// FetchChangeLogWithNotify 带通知的数据获取
func (r *ReaderImpl) FetchChangeLogWithNotify(ctx context.Context) (string, error) {
	select {
	case header := <-NotificationChan:
		// 接收到变更通知，开始处理数据
		return r.FetchChangeLog(ctx, header)
	case <-time.After(1 * time.Second):
		// 超时处理
		return "", fmt.Errorf("timeout waiting for notification")
	case <-ctx.Done():
		// 上下文取消
		return "", ctx.Err()
	}
}

// FetchChangeLog 获取并解析变更日志
func (r *ReaderImpl) FetchChangeLog(ctx context.Context, header ChangLogHeaderInfo) (string, error) {
	// 1. 从数据库获取变更日志记录
	changelog, err := r.fetchChangeLogFromDB(ctx, header)
	if err != nil {
		return "", fmt.Errorf("获取变更日志失败: %w", err)
	}

	// 2. 解析并格式化变更日志
	parsedLog, err := r.ParseChangeLog(changelog)
	if err != nil {
		return "", fmt.Errorf("解析变更日志失败: %w", err)
	}

	return parsedLog, nil
}

// fetchChangeLogFromDB 从数据库获取变更日志记录
func (r *ReaderImpl) fetchChangeLogFromDB(ctx context.Context, header ChangLogHeaderInfo) (*models.TChangeLogs, error) {
	// 1. 解析键名和键值
	keyNames := strings.Split(header.TableKey, ":")
	keyValues := strings.Split(header.TableKeyValue, ":")

	if len(keyNames) != 2 || len(keyValues) != 2 {
		return nil, fmt.Errorf("键格式无效，应为'key1:key2'和'value1:value2'")
	}

	if len(keyNames) != len(keyValues) {
		return nil, fmt.Errorf("键名和键值数量不匹配")
	}

	db, err := mysql.Conn()

	// 2. 构建查询
	query := db.WithContext(ctx).Table(config.TChangeLogs)
	for i := 0; i < len(keyNames); i++ {
		query = query.Where(fmt.Sprintf("%s = ?", keyNames[i]), keyValues[i])
	}

	// 3. 执行查询
	var changelog models.TChangeLogs
	err = query.
		Set("gorm:query_option", "FOR UPDATE"). // 添加行锁
		First(&changelog).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("未找到键为%v，值为%v的变更日志", keyNames, keyValues)
		}
		return nil, fmt.Errorf("数据库查询错误: %w", err)
	}

	return &changelog, nil
}

// ParseChangeLog 解析变更日志为格式化字符串
func (r *ReaderImpl) ParseChangeLog(changelog *models.TChangeLogs) (string, error) {
	if changelog == nil {
		return "", fmt.Errorf("变更日志为空")
	}

	// 解析JSON格式的变更详情
	var changeDetail map[string]interface{}
	if err := json.Unmarshal([]byte(changelog.ChangeLog), &changeDetail); err != nil {
		return "", fmt.Errorf("解析变更详情失败: %w", err)
	}
	return "", nil
}
