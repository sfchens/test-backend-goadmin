package sys_model

import "csf/common/mysql/model"

type MenuListItem struct {
	model.SysMenu
	Children []MenuListItem `json:"children" gorm:"-"`
}
