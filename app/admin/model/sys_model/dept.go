package sys_model

import (
	"time"
)

type DeptTreeListItem struct {
	SysDept
	Children []DeptTreeListItem `json:"children" gorm:"-"`
}

type SysDept struct {
	ID        int       `json:"id"`
	ParentId  int       `json:"parent_id"`
	Label     string    `json:"label"`
	Name      string    `json:"name"`
	Leader    string    `json:"leader"`
	Sort      int       `json:"sort"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Status    int8      `json:"status"`
	Operator  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}
