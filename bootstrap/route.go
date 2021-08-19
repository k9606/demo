package bootstrap

import (
	"demo/routes"
	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute() *gin.Engine {
	router := gin.Default()

	routes.RegisterApiRoutes(router)

	return router
}
