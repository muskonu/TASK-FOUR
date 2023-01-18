package pkg

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func CreateDb() *gorm.DB {
	f, err := os.OpenFile("./log/SQL.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	newLogger := logger.New(
		log.New(f, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,        // 禁用彩色打印
		},
	)
	dsn := "root:12345678@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func LogFile() {
	f, _ := os.OpenFile("./log/gin.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	gin.DefaultWriter = f
}
