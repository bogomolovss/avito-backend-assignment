package model

import "gorm.io/gorm"

// `user_segments` is the join table
type User struct {
	gorm.Model
	Username string    `gorm:"size:255;not null;unique" json:"username"`
	Segment  []Segment `gorm:"many2many:user_segments;"`
}
