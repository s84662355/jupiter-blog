package core

import (
	"github.com/douyu/jupiter/pkg/store/gorm"
)

type Gorm struct {
	*gorm.DB
	config string
}

func (g Gorm) New(config string) *Gorm {
	var gormDB *gorm.DB = gorm.StdConfig(config).Build()
	gg := &Gorm{
		DB:     gormDB,
		config: config,
	}
	return gg
}
