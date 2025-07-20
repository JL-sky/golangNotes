package loger

import (
	"fmt"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	// 设置日志级别
	log.SetLevel(log.DebugLevel)

	// 自定义日志格式
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,                  // 启用完整时间戳
		TimestampFormat: "2006-01-02 15:04:05", // 自定义时间格式
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// 简化文件路径
			path := f.File
			if relPath, err := filepath.Rel(getProjectRoot(), path); err == nil {
				path = relPath
			}

			// 返回文件名和行号
			return "", fmt.Sprintf("%s:%d", path, f.Line)
		},
	})

	// 启用调用者信息记录
	log.SetReportCaller(true)
}

// 获取项目根目录
func getProjectRoot() string {
	_, file, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(file)) // 假设项目根目录是当前文件的上两级目录
	return projectRoot
}
