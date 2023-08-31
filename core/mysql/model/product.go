package model

import (
	"time"
)

type ProductCategory struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"商品分类表ID"`                  // 商品分类表ID
	PID       int       `gorm:"column:pid;not null;default:0" json:"pid" comment:"父id"`                              // 父id
	PIDs      string    `gorm:"column:pids;not null;size:100" json:"pids" comment:"父级分类"`                         // 分类名称
	Name      string    `gorm:"column:name;not null;size:100" json:"name" comment:"分类名称"`                         // 分类名称
	Sort      int       `gorm:"column:sort;not null;default:0" json:"sort" comment:"排序"`                            // 排序
	Pic       string    `gorm:"column:pic;not null;type:text" json:"pic" comment:"图标"`                              // 图标
	BigPic    string    `gorm:"column:big_pic;not null;type:text" json:"big_pic" comment:"分类大图"`                  // 分类大图
	IsShow    int       `gorm:"column:is_show;not null;default:1" json:"is_show" comment:"是否展示"`                  // 是否展示
	Operator  string    `gorm:"column:operator;not null;size:32" json:"operator" comment:"操作人"`                    // 操作人
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"`     // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"最后更新时间"` // 最后更新时间
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at" comment:"删除时间"`                               // 删除时间
}
