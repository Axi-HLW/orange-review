package model

import (
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model
	CreateBy string         `gorm:"size:48;not null;default:'';column:create_by;comment:创建方标识"`
	UpdateBy string         `gorm:"size:48;not null;default:'';column:update_by;comment:更新方标识"`
	Version  uint           `gorm:"not null;default:0;column:version;comment:乐观锁标记"`
}
