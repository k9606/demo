package models

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID        uint           `gorm:"not null;autoIncrement;primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at;index"`
	UpdatedAt time.Time      `gorm:"column:updated_at;index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
