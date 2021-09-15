package middlewares

import (
	"demo/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.ParseToken(c.GetHeader("Token"))

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  err.Error(),
				"data": nil,
			})

			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
