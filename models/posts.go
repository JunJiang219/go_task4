package models

import "gorm.io/gorm"

type Posts struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint
	Comments []Comment
}

func (obj Posts) TableName() string {
	return "posts"
}
