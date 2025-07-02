package main

import (
	"fmt"
	"log"

	changelog "github.com/jl-sky/grom/golangNotes/datatbase/changeLog"
	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	"github.com/jl-sky/grom/golangNotes/datatbase/mysql"
	"gorm.io/gorm"
)

func db_test() {
	db := mysql.Conn()
	defer func() {
		mysql.Close(db)
	}()
	if db == nil {
		return
	}
	//  插入数据
	newVideo := &models.Video{
		Title:       "GORM 教程",
		Description: "学习如何使用 GORM 操作 MySQL",
		URL:         "https://example.com/gorm-tutorial",
	}
	err := mysql.CreateVideo(db, newVideo)
	if err != nil {
		fmt.Println("插入失败:", err)
	} else {
		fmt.Println("插入成功, ID:", newVideo.ID)
	}

	// 查询数据
	video, err := mysql.GetVideoByID(db, newVideo.ID)
	if err != nil {
		fmt.Println("查询失败:", err)
	} else {
		fmt.Printf("查询结果: %+v\n", video)
	}

	// 更新数据
	video.Title = "GORM 高级教程"
	err = mysql.UpdateVideo(db, video)
	if err != nil {
		fmt.Println("更新失败:", err)
	} else {
		fmt.Println("更新成功")
	}

	// 删除数据
	err = mysql.DeleteVideo(db, video)
	if err != nil {
		fmt.Println("删除失败:", err)
	} else {
		fmt.Println("删除成功")
	}
}

func main() {
	db := mysql.Conn() // 获取数据库连接

	// 初始化变更记录系统
	if err := changelog.InitChangeLogSystem(db); err != nil {
		log.Fatal("初始化变更记录系统失败:", err)
	}

	// 创建新视频
	newVideo := &models.Video{
		Title:       "GORM教程",
		Description: "学习GORM基础",
		URL:         "http://example.com/gorm",
	}
	if err := changelog.HandleVideoChange(db, newVideo); err != nil {
		log.Fatal("处理视频变更失败:", err)
	}

	// 更新视频
	updatedVideo := &models.Video{
		Model:       gorm.Model{ID: 1},
		Title:       "GORM高级教程",
		Description: "学习GORM高级特性",
		URL:         "http://example.com/gorm-advanced",
	}
	if err := changelog.HandleVideoChange(db, updatedVideo); err != nil {
		log.Fatal("处理视频变更失败:", err)
	}
}
