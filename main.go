package main

import (
	"demo/bootstrap"
)

func main() {
	// 初始化路由绑定
	router := bootstrap.SetupRoute()

	router.Run(":8080")
}
