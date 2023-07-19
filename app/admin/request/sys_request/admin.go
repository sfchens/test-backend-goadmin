package sys_request

type AdminAddReq struct {
	Username string `json:"username" form:"username" validate:"required" msg:"required:账号必填" description:"账号"`
	Realname string `json:"realname" form:"realname" validate:"required" msg:"required:姓名必填" description:"姓名"`
	Email    string `json:"email" form:"email"  description:"邮箱"`
	Phone    string `json:"phone" form:"phone"  description:"电话号码"`
	Password string `json:"password" form:"password" validate:"required" msg:"required:密码必填"  description:"密码"`
}

type AdminEditReq struct {
	Id       int    `json:"id" form:"id" validate:"required" msg:"required:参数异常" description:"ID"`
	Username string `json:"username" form:"username" validate:"required" msg:"required:账号必填" description:"账号"`
	Realname string `json:"realname" form:"realname" validate:"required" msg:"required:姓名必填" description:"姓名"`
	Email    string `json:"email" form:"email"  description:"邮箱"`
	Phone    string `json:"phone" form:"phone"  description:"电话号码"`
	Password string `json:"password" form:"password" validate:"required" msg:"required:密码必填"  description:"密码"`
}

type AdminSetStatusReq struct {
	Id     int   `json:"id" form:"id" validate:"required" msg:"required:参数异常" description:"ID"`
	Status uint8 `json:"status" form:"status" validate:"required" msg:"required:状态异常" description:"状态"`
}

type AdminListReq struct {
	Page     int    `json:"page" form:"page"  default:"1" description:"页码"`
	PageSize int    `json:"page_size" form:"page_size"  default:"20" description:"页数"`
	Username string `json:"username" form:"username"  description:"名称"`
	Realname string `json:"realname" form:"realname" description:"真名"`
	Email    string `json:"email" form:"email" description:"邮箱"`
	Phone    string `json:"phone" form:"phone" description:"手机号码"`
}

type AdminListRes struct {
	Total int64           `json:"total" form:"total" description:"总数"`
	List  []AdminListItem `json:"list" form:"list" description:"列表"`
}

type AdminListItem struct {
	ID         uint   `gorm:"column:id;primary_key;auto_increment" json:"id"`
	Username   string `gorm:"column:username;type:varchar(16);not null;default:''" json:"username"`
	Realname   string `gorm:"column:realname;type:varchar(16);not null;default:''" json:"realname"`
	Email      string `gorm:"column:email;type:varchar(32);not null;default:''" json:"email"`
	Phone      string `gorm:"column:phone;type:char(11);default:null" json:"phone"`
	HeadPic    string `gorm:"column:head_pic;type:text;not null" json:"head_pic"`
	LastIP     string `gorm:"column:last_ip;type:varchar(16);not null;default:''" json:"last_ip"`
	LastTime   uint   `gorm:"column:last_time;type:int(10);unsigned;not null;default:0" json:"last_time"`
	LoginCount uint   `gorm:"column:login_count;type:int(10);unsigned;not null;default:0" json:"login_count"`
	Status     uint8  `gorm:"column:status;type:tinyint(3);unsigned;not null;default:1" json:"status"`
	Operator   string `gorm:"column:operator;type:varchar(32);not null;default:''" json:"operator"`
	CreatedAt  string `gorm:"column:created_at;not null;default:current_timestamp" json:"created_at"`
	UpdatedAt  string `gorm:"column:updated_at;default:current_timestamp;on update current_timestamp" json:"updated_at"`
}

type AdminInfoRes struct {
	Id           int      `json:"id" form:"id" description:"ID"`
	Username     string   `json:"username" form:"username"  description:"名称"`
	Realname     string   `json:"realname" form:"realname" description:"真名"`
	Email        string   `json:"email" form:"email" description:"邮箱"`
	Phone        string   `json:"phone" form:"phone" description:"手机号码"`
	Roles        []string `json:"roles" form:"roles" description:"手机号码"`
	Permissions  []string `json:"permissions" form:"permissions" description:"权限"`
	Avatar       string   `json:"avatar" form:"avatar" description:"头像"`
	Introduction string   `json:"introduction" form:"introduction" description:"介绍"`
}
