package routes

import (
	"demo/app/http/controllers"
	"demo/app/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine) {
	r.POST("/login", new(controllers.DemoController).Login)

	demo := r.Group("/demo").Use(middlewares.Jwt())
	{
		dc := new(controllers.DemoController)
		demo.POST("/create", dc.Create)
		demo.POST("/delete", dc.Delete)
		demo.POST("/read", dc.Read)
		demo.POST("/update", dc.Update)
		demo.POST("/list", dc.List)
	}
}
