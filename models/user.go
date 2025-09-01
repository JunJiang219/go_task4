package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;type:varchar(255) not null" form:"username"`
	Password string `gorm:"type:varchar(100) not null" form:"password"`
	Email    string `gorm:"type:varchar(255) not null" form:"email"`
	Posts    []Posts
}
