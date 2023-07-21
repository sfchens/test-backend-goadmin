package sys_model

import "csf/common/mysql/model"

type DeptTreeListItem struct {
	model.SysDept
	Children []DeptTreeListItem `json:"children" gorm:"-"`
}
