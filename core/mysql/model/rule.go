package model

import "time"

type ProductRule struct {
	ID        int       `gorm:"primaryKey" json:"id" comment:"自增ID"`
	Name      string    `gorm:"column:name" json:"name" comment:"规格名称"`
	Value     string    `gorm:"type:text" json:"value" comment:"规格值"`
	Operator  string    `gorm:"column:operator" json:"operator" comment:"操作人"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}
