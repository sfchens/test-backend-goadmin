package model

import (
	"time"
)

type SysAdmin struct {
	ID         uint      `gorm:"column:id;primaryKey" json:"id"`                                // ID
	Username   string    `gorm:"column:username;size:16;not null" json:"username"`              // 用户名
	Realname   string    `gorm:"column:realname;size:16;not null" json:"realname"`              // 真实姓名
	Email      string    `gorm:"column:email;size:32;not null" json:"email"`                    // 邮箱
	Phone      string    `gorm:"column:phone;size:11" json:"phone,omitempty"`                   // 电话
	HeadPic    string    `gorm:"column:head_pic;type:text;not null" json:"head_pic"`            // 头像
	Password   string    `gorm:"column:password;size:100;not null" json:"password,omitempty"`   // 密码
	LastIP     string    `gorm:"column:last_ip;size:16;not null" json:"last_ip"`                // 最后登录IP
	LastTime   int       `gorm:"column:last_time;default:0" json:"last_time"`                   // 最后登录时间
	LoginCount int       `gorm:"column:login_count;default:0" json:"login_count"`               // 登录次数
	Status     int       `gorm:"column:status;default:1" json:"status"`                         // 状态
	Operator   string    `gorm:"column:operator;size:32;not null" json:"operator"`              // 操作人
	RoleIds    string    `gorm:"column:role_ids;size:255;not null" json:"operator"`             // 操作人
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
	DeptID     int       `gorm:"column:dept_id;default:0" json:"dept_id"`                       // 部门ID
	Sex        int       `gorm:"column:sex;default:1" json:"sex"`                               // 性别
	Remark     string    `gorm:"column:remark;size:100;not null" json:"remark"`                 // 备注
}

type SysConfig struct {
	ID        uint      `gorm:"column:id;primary_key" json:"id" comment:"ID"`
	Name      string    `gorm:"column:name" json:"name" comment:"名称"`
	Key       string    `gorm:"column:key" json:"key" comment:"0json配置1基础配置2商城配置3用户配置"`
	Config    string    `gorm:"column:config" json:"config" comment:"配置"`
	IsOpen    uint      `gorm:"column:is_open" json:"is_open" comment:"是否开启"`
	Remark    string    `gorm:"column:remark" json:"remark" comment:"备注"`
	Operator  string    `gorm:"column:operator" json:"operator" comment:"操作人"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" comment:"更新时间"`
}

type SysUser struct {
	ID         uint      `gorm:"column:id;primary_key;auto_increment" json:"id"`
	Username   string    `gorm:"column:username" json:"username"`
	Realname   string    `gorm:"column:realname" json:"realname"`
	Email      string    `gorm:"column:email" json:"email"`
	Phone      string    `gorm:"column:phone" json:"phone"`
	Type       int       `gorm:"column:type" json:"type"`
	HeadPic    string    `gorm:"column:head_pic" json:"head_pic"`
	Password   string    `gorm:"column:password" json:"-"`
	LastIP     string    `gorm:"column:last_ip" json:"-"`
	LastTime   int       `gorm:"column:last_time" json:"last_time"`
	LoginCount int       `gorm:"column:login_count" json:"login_count"`
	Status     int       `gorm:"column:status" json:"status"`
	Operator   string    `gorm:"column:operator" json:"operator"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type SysMenu struct {
	Id         int        `gorm:"AUTO_INCREMENT;column:id" json:"id"`                                                        // ID
	MenuName   string     `gorm:"column:menu_name" json:"menu_name"`                                                         // 菜单名称
	Title      string     `gorm:"column:title" json:"title"`                                                                 // 标题
	Icon       string     `gorm:"column:icon" json:"icon"`                                                                   // 图标
	Path       string     `gorm:"column:path" json:"path"`                                                                   // 前端路径
	ParentId   int        `gorm:"column:parent_id" json:"parent_id"`                                                         // 父级
	ParentIds  string     `gorm:"column:parent_ids" json:"parent_ids"`                                                       // 父级类型
	MenuType   string     `gorm:"column:menu_type" json:"menu_type"`                                                         // 菜单类型，M目录 C菜单，F按钮
	Permission string     `gorm:"column:permission" json:"permission"`                                                       // 权限标识
	Component  string     `gorm:"column:component" json:"component"`                                                         // 组件
	ApisId     string     `gorm:"column:apis_id" json:"apis_id"`                                                             // 组件
	Sort       int        `gorm:"column:sort" json:"sort"`                                                                   // 排序
	Visible    int        `gorm:"column:visible" json:"visible"`                                                             // 是否启用，1启用
	IsFrame    int        `gorm:"column:is_frame" json:"is_frame"`                                                           // 是否框架，1
	Operator   string     `gorm:"NOT NULL;DEFAULT:'';column:operator" json:"operator"`                                       // 操作人
	CreatedAt  time.Time  `gorm:"DEFAULT:current_timestamp;column:created_at" json:"created_at"`                             // 创建时间
	UpdatedAt  time.Time  `gorm:"DEFAULT:current_timestamp ON UPDATE current_timestamp;column:updated_at" json:"updated_at"` // 最后更新时间
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at"`                                                       // 删除时间
}

type SysMenuApiRule struct {
	SysMenuID uint64 `gorm:"column:sys_menu_id;uniqueIndex:menu_id_api_id_key" json:"sys_menu_id" comment:"菜单ID"`
	SysApiID  uint64 `gorm:"column:sys_api_id;uniqueIndex:menu_id_api_id_key" json:"sys_api_id" comment:"api id"`
}

type SysDept struct {
	ID        int       `gorm:"column:id;primary_key;auto_increment" json:"id"`
	ParentId  int       `gorm:"column:parent_id" json:"parent_id"`
	Name      string    `gorm:"column:name" json:"name"`
	Leader    string    `gorm:"column:leader" json:"leader"`
	Sort      int       `gorm:"column:sort" json:"sort"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	Status    int8      `gorm:"column:status" json:"status"`
	Operator  string    `gorm:"column:operator;not null;default:''" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"-"`
}

type SysSwitch struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Key       string    `gorm:"column:key" json:"key"`
	Status    int       `gorm:"column:status" json:"status"`
	Remark    string    `gorm:"column:remark" json:"remark"`
	Operator  string    `gorm:"column:operator" json:"operator"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}

type SysApi struct {
	ID        uint64     `gorm:"column:id;primaryKey" json:"id"`
	Tags      string     `gorm:"column:tags" json:"tags"`
	Title     string     `gorm:"column:title" json:"title"`
	Path      string     `gorm:"column:path" json:"path"`
	Method    string     `gorm:"column:method" json:"method"`
	Handle    string     `gorm:"column:handle" json:"handle"`
	Operator  string     `gorm:"column:operator" json:"operator"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type SysRole struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"表主键"`
	Name      string    `gorm:"column:name;size:128" json:"name" comment:"名称"`
	Status    int       `gorm:"column:status;type:tinyint" json:"status" comment:"状态"`
	Key       string    `gorm:"column:key;size:128" json:"key" comment:"权限标识"`
	Sort      int64     `gorm:"column:sort" json:"sort" comment:"排序"`
	MenuIds   string    `gorm:"column:menu_ids" json:"menu_ids" comment:"菜单ID"`
	Remark    string    `gorm:"column:remark;size:255" json:"remark" comment:"备注"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at" comment:"最后更新时间"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at" comment:"删除时间"`
}

type SysRoleMenu struct {
	RoleID int64 `gorm:"column:role_id;primaryKey" json:"role_id" comment:"角色ID"`
	MenuID int64 `gorm:"column:menu_id;primaryKey" json:"menu_id" comment:"菜单ID"`
}
