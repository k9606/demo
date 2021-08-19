package controllers

import (
	"github.com/gin-gonic/gin"
)

type DemoController struct{}

func (s *DemoController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create",
	})
}

func (s *DemoController) Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete",
	})
}

func (s *DemoController) Read(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Read",
	})
}

func (s *DemoController) Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update",
	})
}

func (s *DemoController) List(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List",
	})
}
