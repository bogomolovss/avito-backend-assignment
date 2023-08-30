package model

import (
	"user-segments/database"

	"gorm.io/gorm"
)

type Segment struct {
	gorm.Model
	Title string  `gorm:"type:text" json:"title"`
	Users []*User `gorm:"many2many:user_segments;"`
}

func (segment *Segment) Save() (*Segment, error) {
	err := database.Database.Create(&segment).Error
	if err != nil {
		return &Segment{}, err
	}
	return segment, nil
}
