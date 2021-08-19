package routes

import (
	"demo/app/http/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterApiRoutes 注册网页相关路由
func RegisterApiRoutes(r *gin.Engine) {
	demo := r.Group("/demo")
	{
		demo.POST("/create", (&controllers.DemoController{}).Create)
		demo.POST("/delete", (&controllers.DemoController{}).Delete)
		demo.POST("/read", (&controllers.DemoController{}).Read)
		demo.POST("/update", (&controllers.DemoController{}).Update)
		demo.POST("/list", (&controllers.DemoController{}).List)
	}
}
