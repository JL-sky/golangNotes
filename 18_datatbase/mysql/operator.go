package mysql

import (
	"github.com/jl-sky/grom/golangNotes/datatbase/consts"
	"github.com/jl-sky/grom/golangNotes/datatbase/models"
	"gorm.io/gorm"
)

// 插入数据，向数据库插入一条 Video 记录
func CreateVideo(db *gorm.DB, video *models.Video) error {
	result := db.Table(consts.TableName).Create(video)
	return result.Error
}

// 查询数据，根据 id 查询 Video 记录
func GetVideoByID(db *gorm.DB, id uint) (*models.Video, error) {
	var video models.Video
	result := db.Table(consts.TableName).First(&video, id)
	return &video, result.Error
}

// 更新数据，更新 Video 记录（基于 video.ID 更新）
func UpdateVideo(db *gorm.DB, video *models.Video) error {
	result := db.Table(consts.TableName).Save(video)
	return result.Error
}

// 删除数据，删除 Video 记录（基于 video.ID 删除）
func DeleteVideo(db *gorm.DB, video *models.Video) error {
	result := db.Table(consts.TableName).Delete(video)
	return result.Error
}
