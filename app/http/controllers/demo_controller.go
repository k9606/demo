package controllers

import (
	"demo/app/models/demo"
	"demo/pkg/jwt"
	"demo/pkg/logger"
	"demo/pkg/model"
	"demo/pkg/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DemoController struct{}

func (s *DemoController) Login(c *gin.Context) {
	token, err := jwt.GenerateToken()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  logger.LogError(err),
		"data": gin.H{
			"token": token,
		},
	})
}

func (s *DemoController) Create(c *gin.Context) {
	m := demo.Demo{}
	c.ShouldBind(&m)

	model.DB.Create(&m)

	c.JSON(200, gin.H{
		"message": "Create",
	})
}

func (s *DemoController) Delete(c *gin.Context) {
	m := demo.Demo{}
	c.ShouldBind(&m)

	model.DB.Delete(&m)

	c.JSON(200, gin.H{
		"message": "Delete",
	})
}

func (s *DemoController) Read(c *gin.Context) {
	m := demo.Demo{}
	c.ShouldBind(&m)

	model.DB.First(&m)

	c.JSON(200, gin.H{
		"message": m,
	})
}

func (s *DemoController) Update(c *gin.Context) {
	m := demo.Demo{}
	c.ShouldBind(&m)

	model.DB.Updates(&m)

	c.JSON(200, gin.H{
		"message": "Update",
	})
}

func (s *DemoController) List(c *gin.Context) {
	param := pagination.Param{}
	c.ShouldBind(&param)

	articles := make([]*demo.Demo, 0)

	db := model.DB.Table("demos")

	db, count := pagination.Paginate(db, param)

	db = db.Order("id desc").
		Find(&articles)

	c.JSON(200, gin.H{
		"message": articles,
		"count":   count,
	})
}
