package model

import "time"

type LiveBackdrop struct {
	ID        uint      `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Type      int       `gorm:"column:type" json:"type"`
	Status    int       `gorm:"column:status" json:"status"`
	URL       string    `gorm:"column:url" json:"url"`
	Operator  string    `gorm:"column:operator" json:"operator"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type LiveVideo struct {
	ID        uint      `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Type      int       `gorm:"column:type" json:"type"`
	Status    int       `gorm:"column:status" json:"status"`
	URL       string    `gorm:"column:url" json:"url"`
	Operator  string    `gorm:"column:operator" json:"operator"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
