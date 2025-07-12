package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conn() (*gorm.DB, error) {
	dsn := "root:54264534@tcp(192.168.190.133:3306)/mvldata?charset=utf8mb4&parseTime=True&loc=Local"

	// 1. 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return nil, err
	}
	fmt.Println("连接数据库成功")

	// 2. 获取底层 SQL 数据库对象并配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("获取底层数据库连接失败：", err)
		return nil, err
	}

	// 连接池配置
	sqlDB.SetMaxIdleConns(10)           // 空闲连接池的最大连接数
	sqlDB.SetMaxOpenConns(100)          // 数据库的最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接可复用的最大时间
	return db, nil
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
