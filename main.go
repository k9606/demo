package main

import (
	"demo/bootstrap"
	"demo/config"
	c "demo/pkg/config"
	"demo/pkg/logger"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()
	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	err := router.Run(":" + c.GetString("app.port"))
	logger.LogError(err)
}
