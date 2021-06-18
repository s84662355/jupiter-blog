package models

import (
	"JupiterBlog/core"
)

var db *core.Gorm

func init() {
	db = core.Gorm{}.New("default")
}

func GetDb() *core.Gorm {
	return db
}
