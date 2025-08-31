package models

import "gorm.io/gorm"

type Posts struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255) not null"`
	Content  string `gorm:"not null"`
	UserID   uint
	Comments []Comment
}

// func (obj Posts) TableName() string {
// 	return "posts"
// }
