package models

import "gorm.io/gorm"

type Posts struct {
	gorm.Model
	Title    string    `gorm:"type:varchar(255) not null" json:"title"`
	Content  string    `gorm:"not null" json:"content"`
	UserID   uint      `json:"user_id"`
	Comments []Comment `json:"comments"`
}

func (obj Posts) TableName() string {
	return "posts"
}
