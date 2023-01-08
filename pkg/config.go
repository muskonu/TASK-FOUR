package pkg

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDb() *gorm.DB {
	dsn := "root:fhs123456.@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
