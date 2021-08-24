package routes

import (
	"demo/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {
	demo := r.Group("/demo")
	{
		dc := new(controllers.DemoController)
		demo.POST("/login", dc.Login)
		demo.POST("/create", dc.Create)
		demo.POST("/delete", dc.Delete)
		demo.POST("/read", dc.Read)
		demo.POST("/update", dc.Update)
		demo.POST("/list", dc.List)
	}
}
