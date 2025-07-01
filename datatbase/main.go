package main

import (
	"fmt"

	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	"github.com/jl-sky/grom/golangNotes/datatbase/mysql"
)

func main() {
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
	video, err := mysql.GetVideoByID(db, 1) // 假设 ID=1 的记录存在
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
