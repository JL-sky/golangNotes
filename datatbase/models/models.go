package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Title       string
	Description string
	URL         string
}
