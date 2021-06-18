package models

import (
	"JupiterBlog/lib/helper"
	"github.com/douyu/jupiter/pkg/store/gorm"
)

// 文章分类表
type Cate struct {
	Id        uint32           `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string           `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`
	CreatedAt *helper.JSONTime `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt *helper.JSONTime `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	DeletedAt *helper.JSONTime `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
}

func (m Cate) Model() *gorm.DB {
	return db.Model(&m)
}

func (m Cate) TableName() string {
	return "cate"
}
