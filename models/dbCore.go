package models

import (
	"sync"

	"github.com/go_task4/utils"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(func() {
		dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			utils.GetLogger().Error(
				"数据库连接出错",
				zap.String("error", err.Error()),
			)
			return
		}
	})
	return db
}

func AutoMigrate() {
	GetDB().AutoMigrate(&User{}, &Posts{}, &Comment{})
}
