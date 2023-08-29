package model

import "gorm.io/gorm"

type Segment struct {
	gorm.Model
	Title  string `gorm:"type:text" json:"title"`
	UserID uint
}
