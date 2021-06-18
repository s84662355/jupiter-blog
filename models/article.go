package models

import (
	"JupiterBlog/lib/helper"
	"github.com/douyu/jupiter/pkg/store/gorm"
)

// 文章表
type Article struct {
	Id         uint32           `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Title      string           `gorm:"column:title;type:varchar(100);NOT NULL" json:"title"`
	Content    string           `gorm:"column:content;type:longtext;NOT NULL" json:"content"`
	CateId     uint32           `gorm:"column:cate_id;type:int(11) unsigned;default:0;NOT NULL" json:"cate_id"`
	ReadAmount uint64           `gorm:"column:read_amount;type:bigint(20) unsigned;default:0;NOT NULL" json:"read_amount"`
	Sort       uint             `gorm:"column:sort;type:int(11) unsigned;default:0;NOT NULL" json:"sort"`
	CreatedAt  *helper.JSONTime `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt  *helper.JSONTime `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	DeletedAt  *helper.JSONTime `gorm:"column:deleted_at;type:datetime" json:"deleted_at"`
	Status     uint             `gorm:"column:status;type:int(11) unsigned;default:1;NOT NULL" json:"status"`
	Image      string           `gorm:"column:image;type:varchar(200);NOT NULL" json:"image"`
	Summary    string           `gorm:"column:summary;type:varchar(500);NOT NULL" json:"summary"`
}

func (m Article) Model() *gorm.DB {
	return db.Model(&m)
}

func (m Article) TableName() string {
	return "article"
}
