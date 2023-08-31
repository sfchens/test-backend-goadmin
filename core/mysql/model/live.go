package model

import "time"

type LiveBackdrop struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"ID"`                             // ID
	Name      string    `gorm:"column:name;not null;default:''" json:"name" comment:"名称"`                              // 名称
	Type      int       `gorm:"column:type;default:0" json:"type" comment:"图片分类"`                                      // 图片分类
	Status    int       `gorm:"column:status;default:0" json:"status" comment:"启用状态，0未开启1开启"`                          // 启用状态
	URL       string    `gorm:"column:url" json:"url" comment:"图片地址"`                                                  // 图片地址
	Operator  string    `gorm:"column:operator;not null;default:''" json:"operator" comment:"操作人"`                     // 操作人
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"更新时间"`          // 更新时间
}
type LiveVideo struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"ID"`                             // ID
	Name      string    `gorm:"column:name;not null;default:''" json:"name" comment:"名称"`                              // 名称
	Type      int       `gorm:"column:type;default:0" json:"type" comment:"视频分类"`                                      // 视频分类
	Status    int       `gorm:"column:status;default:0" json:"status" comment:"启用状态，0未开启1开启"`                          // 启用状态
	URL       string    `gorm:"column:url" json:"url" comment:"图片地址"`                                                  // 图片地址
	Operator  string    `gorm:"column:operator;not null;default:''" json:"operator" comment:"操作人"`                     // 操作人
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"更新时间"`          // 更新时间
}
