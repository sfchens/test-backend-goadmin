package model

import (
	"time"
)

type SysAdmin struct {
	ID         uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"ID"`                             // ID
	Username   string     `gorm:"column:username;not null;default:''" json:"username" comment:"账号"`                      // 账号
	Realname   string     `gorm:"column:realname;not null;default:''" json:"realname" comment:"姓名"`                      // 姓名
	Email      string     `gorm:"column:email;not null;default:''" json:"email" comment:"邮箱"`                            // 邮箱
	Phone      string     `gorm:"column:phone;default:''" json:"phone" comment:"手机号码"`                                   // 手机号码
	HeadPic    string     `gorm:"column:head_pic;not null" json:"head_pic" comment:"头像"`                                 // 头像
	Password   string     `gorm:"column:password;not null;default:''" json:"password" comment:"密码"`                      // 密码
	LastIP     string     `gorm:"column:last_ip;not null;default:''" json:"last_ip" comment:"最后登录IP"`                    // 最后登录IP
	LastTime   *time.Time `gorm:"column:last_time" json:"last_time" comment:"最后登录时间"`                                    // 最后登录时间
	LoginCount int        `gorm:"column:login_count;not null;default:0" json:"login_count" comment:"登录次数"`               // 登录次数
	Status     int        `gorm:"column:status;not null;default:1" json:"status" comment:"状态，0未激活 1已激活 2已禁用"`            // 状态
	DeptID     int        `gorm:"column:dept_id;not null;default:0" json:"dept_id" comment:"部门ID"`                       // 部门ID
	RoleIDS    string     `gorm:"column:role_id;not null;default:0" json:"role_id" comment:"角色ID"`                       // 角色ID
	Sex        int        `gorm:"column:sex;not null;default:0" json:"sex" comment:"性别，0未知 1男 2女"`                       // 性别
	Remark     string     `gorm:"column:remark;not null;default:''" json:"remark" comment:"备注"`                          // 备注
	Operator   string     `gorm:"column:operator;not null;default:''" json:"operator" comment:"操作人"`                     // 操作人
	CreatedAt  time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"` // 创建时间
	UpdatedAt  time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"更新时间"`          // 更新时间
}

type SysConfig struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"ID"`                             // ID
	Key       string    `gorm:"column:key;not null;default:''" json:"key" comment:"名称"`                                // 名称
	Name      string    `gorm:"column:name;not null;default:''" json:"name" comment:"名称"`                              // 名称
	Config    string    `gorm:"column:config" json:"config" comment:"配置"`                                              // 配置
	IsOpen    int       `gorm:"column:is_open;default:0" json:"is_open" comment:"是否开启"`                                // 是否开启
	Remark    string    `gorm:"column:remark;not null;default:''" json:"remark" comment:"备注"`                          // 备注
	Operator  string    `gorm:"column:operator;not null;default:''" json:"operator" comment:"操作人"`                     // 操作人
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"更新时间"`          // 更新时间
}

type SysMenu struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"ID"`                      // ID
	MenuName   string     `gorm:"column:menu_name" json:"menu_name" comment:"菜单名称"`                               // 菜单名称
	Title      string     `gorm:"column:title" json:"title" comment:"标题"`                                         // 标题
	Icon       string     `gorm:"column:icon" json:"icon" comment:"图标"`                                           // 图标
	Path       string     `gorm:"column:path" json:"path" comment:"前端路径"`                                         // 前端路径
	ParentID   int64      `gorm:"column:parent_id" json:"parent_id" comment:"父级"`                                 // 父级
	ParentIDs  string     `gorm:"column:parent_ids" json:"parent_ids" comment:"父级类型"`                             // 父级类型
	MenuType   string     `gorm:"column:menu_type" json:"menu_type" comment:"菜单类型，M目录 C菜单，F按钮"`                   // 菜单类型
	Permission string     `gorm:"column:permission" json:"permission" comment:"权限标识"`                             // 权限标识
	Component  string     `gorm:"column:component" json:"component" comment:"组件"`                                 // 组件
	Sort       int        `gorm:"column:sort" json:"sort" comment:"排序"`                                           // 排序
	Visible    int        `gorm:"column:visible" json:"visible" comment:"是否启用，1启用"`                               // 是否启用
	IsFrame    int        `gorm:"column:is_frame" json:"is_frame" comment:"是否框架，1"`                               // 是否框架
	Operator   string     `gorm:"column:operator;not null;default:''" json:"operator" comment:"操作人"`              // 操作人
	CreatedAt  time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"`   // 创建时间
	UpdatedAt  time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"最后更新时间"` // 最后更新时间
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"deleted_at" comment:"删除时间"`                             // 删除时间
	ApisID     string     `gorm:"column:apis_id" json:"apis_id" comment:"组件"`                                     // 组件
}

type SysUser struct {
	ID         uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"ID"`                    // ID
	Username   string    `gorm:"column:username;not null;size:16" json:"username" comment:"账号"`                // 账号
	Realname   string    `gorm:"column:realname;not null;size:16" json:"realname" comment:"姓名"`                // 姓名
	Email      string    `gorm:"column:email;not null;size:32" json:"email" comment:"邮箱"`                      // 邮箱
	Phone      string    `gorm:"column:phone;size:11" json:"phone" comment:"手机号码"`                             // 手机号码
	HeadPic    string    `gorm:"column:head_pic;not null" json:"head_pic" comment:"头像"`                        // 头像
	Password   string    `gorm:"column:password;not null;size:100" json:"password" comment:"密码"`               // 密码
	LastIP     string    `gorm:"column:last_ip;not null;size:16" json:"last_ip" comment:"最后登录IP"`              // 最后登录IP
	LastTime   time.Time `gorm:"column:last_time" json:"last_time" comment:"最后登录时间"`                           // 最后登录时间
	LoginCount int       `gorm:"column:login_count;not null" json:"login_count" comment:"登录次数"`                // 登录次数
	Status     int       `gorm:"column:status;not null" json:"status" comment:"状态"`                            // 状态，0未激活 1已激活 2已禁用
	Sex        int       `gorm:"column:sex;not null" json:"sex" comment:"性别"`                                  // 性别， 0未知 1男 2女
	Remark     string    `gorm:"column:remark;not null;size:100" json:"remark" comment:"备注"`                   // 备注
	Operator   string    `gorm:"column:operator;not null;size:32" json:"operator" comment:"操作人"`               // 操作人
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"` // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"更新时间"` // 更新时间
}

type SysMenuAPIRule struct {
	SysMenuID int64 `gorm:"column:sys_menu_id;primaryKey" json:"sys_menu_id" comment:"菜单ID"` // 菜单ID
	SysAPIID  int64 `gorm:"column:sys_api_id;primaryKey" json:"sys_api_id" comment:"api id"` // api id
}

type SysDept struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"ID"`                      // ID
	ParentID  int64      `gorm:"column:parent_id" json:"parent_id" comment:"上级部门"`                               // 上级部门
	Name      string     `gorm:"column:name" json:"name" comment:"部门名称"`                                         // 部门名称
	Leader    string     `gorm:"column:leader" json:"leader" comment:"负责人"`                                      // 负责人
	Sort      int        `gorm:"column:sort" json:"sort" comment:"排序"`                                           // 排序
	Phone     string     `gorm:"column:phone" json:"phone" comment:"手机号码"`                                       // 手机号码
	Email     string     `gorm:"column:email" json:"email" comment:"邮箱"`                                         // 邮箱
	Status    int        `gorm:"column:status" json:"status" comment:"状态"`                                       // 状态
	Operator  string     `gorm:"column:operator;not null;default:''" json:"operator" comment:"操作人"`              // 操作人
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"`   // 创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"最后更新时间"` // 最后更新时间
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" comment:"删除时间"`                             // 删除时间
}

type SysSwitch struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"ID"`                    // ID
	Name      string    `gorm:"column:name" json:"name" comment:"名称"`                                         // 名称
	Key       string    `gorm:"column:key" json:"key" comment:"键名"`                                           // 键名
	Status    int       `gorm:"column:status" json:"status" comment:"状态"`                                     // 状态
	Remark    string    `gorm:"column:remark" json:"remark" comment:"描述"`                                     // 描述
	Operator  string    `gorm:"column:operator" json:"operator" comment:"操作人"`                                // 操作人
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"更新时间"` // 更新时间
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at" comment:"删除时间"`                           // 删除时间
}

type SysAPI struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"ID"`                             // ID
	Tags      string     `gorm:"column:tags" json:"tags" comment:"标题"`                                                  // 标题
	Title     string     `gorm:"column:title" json:"title" comment:"标题"`                                                // 标题
	Path      string     `gorm:"column:path" json:"path" comment:"地址"`                                                  // 地址
	Method    string     `gorm:"column:method" json:"method" comment:"请求类型"`                                            // 请求类型
	Handle    string     `gorm:"column:handle" json:"handle" comment:"handle"`                                          // handle
	Operator  string     `gorm:"column:operator;not null;default:''" json:"operator" comment:"操作人"`                     // 操作人
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"` // 创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"更新时间"`          // 更新时间
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" comment:"删除时间"`                                    // 删除时间
}

type SysRole struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id" comment:"ID"`                      // ID
	Name      string     `gorm:"column:name" json:"name" comment:"名称"`                                           // 名称
	Status    int        `gorm:"column:status" json:"status" comment:"状态"`                                       // 状态
	Key       string     `gorm:"column:key" json:"key" comment:"权限标识"`                                           // 权限标识
	Sort      int64      `gorm:"column:sort" json:"sort" comment:"排序"`                                           // 排序
	Remark    string     `gorm:"column:remark" json:"remark" comment:"备注"`                                       // 备注
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at" comment:"创建时间"`   // 创建时间
	UpdatedAt time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at" comment:"最后更新时间"` // 最后更新时间
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" comment:"删除时间"`                             // 删除时间
	MenuIDs   string     `gorm:"column:menu_ids" json:"menu_ids" comment:"排序"`                                   // 排序
}

type SysRoleMenu struct {
	RoleID int64 `gorm:"column:role_id;primaryKey" json:"role_id" comment:"角色ID"` // 角色ID
	MenuID int64 `gorm:"column:menu_id;primaryKey" json:"menu_id" comment:"菜单ID"` // 菜单ID
}
