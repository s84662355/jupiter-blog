package models

import (
	"github.com/douyu/jupiter/pkg/store/gorm"
)

// 配置表
type Setting struct {
	Id      uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	KeyName string `gorm:"column:key_name;type:varchar(100);NOT NULL" json:"key_name"`
	Content string `gorm:"column:content;type:text;NOT NULL" json:"content"`
}

func (m Setting) Model() *gorm.DB {
	return db.Model(&m)
}

func (m Setting) TableName() string {
	return "setting"
}
