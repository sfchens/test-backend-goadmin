package model

import (
	"gorm.io/gorm"
	"time"
)

type ComPicture struct {
	ID        int            `gorm:"column:id;primaryKey" json:"id"`
	Filename  string         `gorm:"column:filename" json:"filename"`
	Path      string         `gorm:"column:path" json:"path"`
	MD5Str    string         `gorm:"column:md5_str" json:"md5_str"`
	Type      int            `gorm:"column:type" json:"type"`
	Operator  string         `gorm:"column:operator" json:"operator"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}
