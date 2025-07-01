package mysql

import (
	"fmt"
	"time"

	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conn() *gorm.DB {
	dsn := "root:54264534@tcp(192.168.190.132:3306)/grom_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 1. 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return nil
	}
	fmt.Println("连接数据库成功")

	// 2. 获取底层 SQL 数据库对象并配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("获取底层数据库连接失败：", err)
		return nil
	}

	// 连接池配置
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池的最大连接数
	sqlDB.SetMaxOpenConns(100)          // 数据库的最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间

	// 3. 自动迁移（创建表）
	err = db.AutoMigrate(&models.Video{})
	if err != nil {
		fmt.Println("自动迁移失败：", err)
		return nil
	}
	fmt.Println("自动迁移成功")

	return db
}

func Close(db *gorm.DB) {
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			fmt.Println("获取底层数据库连接失败：", err)
			return
		}
		sqlDB.Close() // 关闭数据库连接
		fmt.Println("数据库连接已关闭")
	}()
}
