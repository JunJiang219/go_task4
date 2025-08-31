package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255) not null"`
	Password string `gorm:"type:varchar(100) not null"`
	Email    string `gorm:"type:varchar(255) not null"`
	Posts    []Posts
}
