package demo

import (
	"demo/app/models"
)

// Demo 用户模型
type Demo struct {
	models.BaseModel
	Name  string `gorm:"column:name;type:varchar(255);not null;" valid:"name"`
	Email string `gorm:"column:email;type:varchar(255);default:NULL;" valid:"email"`
}
