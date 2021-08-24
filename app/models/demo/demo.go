package demo

import (
	"demo/app/models"
)

// Demo 用户模型
type Demo struct {
	models.BaseModel
	Name     string `gorm:"type:varchar(20);not null;default:'';comment:姓名"`
	Password string `gorm:"type:varchar(20);not null;default:'';comment:密码"`
	Gender   uint8  `gorm:"type:tinyint(4);not null;default:0;comment:性别 1男 2女;index"`
}
