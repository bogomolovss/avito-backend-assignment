package model

import (
	"html"
	"strings"
	"user-segments/database"

	"gorm.io/gorm"
)

// `user_segments` is the join table
type User struct {
	gorm.Model
	Username string    `gorm:"size:255;not null;unique" json:"username"`
	Segments []Segment `gorm:"many2many:user_segments;"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// returns error if we want add password in future
func (user *User) BeforeSave(*gorm.DB) error {
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

// Retrieve user list with eager loading segments
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("Segments").Find(&users).Error
	return users, err
}
