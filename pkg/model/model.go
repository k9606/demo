package model

import (
	"demo/pkg/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB gorm.DB 对象
var DB *gorm.DB

// ConnectDB 初始化模型
func ConnectDB() *gorm.DB {

	var err error

	// 初始化 MySQL 连接信息
	var (
		host     = "127.0.0.1"
		port     = "3306"
		database = "demo"
		username = "root"
		password = "qqqqqq"
		charset  = "utf8mb4"
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")

	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	//var level gormlogger.LogLevel
	//if config.GetBool("app.debug") {
	//	// 读取不到数据也会显示
	//	level = gormlogger.Warn
	//} else {
	//	// 只有错误才会显示
	//	level = gormlogger.Error
	//}

	// 准备数据库连接池
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		//Logger: gormlogger.Default.LogMode(level),
	})

	logger.LogError(err)

	return DB
}
