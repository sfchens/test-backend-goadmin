package sys_service

import (
	"csf/app/admin/request/sys_request"
	"csf/common/mysql/model"
	"csf/library/db"
	"csf/library/my_jwt"
	"csf/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sSysAdmin struct {
	ctx *gin.Context
}

func NewSysAdminService(ctx *gin.Context) *sSysAdmin {
	return &sSysAdmin{ctx: ctx}
}

func (s *sSysAdmin) Add(input sys_request.AdminAddReq) (err error) {
	var (
		username = input.Username
		realname = input.Realname
		password = input.Password
		phone    = input.Phone
		email    = input.Email
		remark   = input.Remark
		sex      = input.Sex
		deptId   = input.DeptId
		roleIds  = input.RoleIds
		status   = input.Status

		sysAdminModel model.SysAdmin
	)
	fmt.Printf("roleIds:  %+v\n", roleIds)
	var counts int64
	err = db.GetDb().Model(sysAdminModel).Where("username=?", username).Count(&counts).Error
	if err != nil {
		return
	}
	if counts > 0 {
		err = errors.New("该账号已存在")
		return
	}

	if phone != "" {
		var phoneExistCount int64
		err = db.GetDb().Model(sysAdminModel).Where("phone = ?", phone).Count(&phoneExistCount).Error
		if err != nil {
			return
		}

		if phoneExistCount > 0 {
			err = errors.New("该号码已被绑定")
			return
		}
		sysAdminModel.Phone = phone
	}

	if email != "" {
		var emailExistCount int64
		err = db.GetDb().Model(sysAdminModel).Where("email = ?", email).Count(&emailExistCount).Error
		if err != nil {
			return
		}

		if emailExistCount > 0 {
			err = errors.New("该邮箱已被绑定")
			return
		}

		sysAdminModel.Email = email
	}

	// admin表
	sysAdminModel.Remark = remark
	sysAdminModel.Sex = sex
	sysAdminModel.DeptID = deptId
	sysAdminModel.Status = status
	sysAdminModel.Username = username
	sysAdminModel.Realname = realname
	sysAdminModel.Password = utils.BcryptHash(password)

	// 事务更新数据
	err = db.GetDb().Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&sysAdminModel).Error
		if err != nil {
			return
		}
		return
	})
	if err != nil {
		return
	}

	return
}

func (s *sSysAdmin) Edit(input sys_request.AdminEditReq) (err error) {
	var (
		id       = input.Id
		realname = input.Realname
		email    = input.Email
		password = input.Password
		phone    = input.Phone

		sysAdminModel model.SysAdmin
	)

	err = db.GetDb().First(&sysAdminModel, id).Error
	if err != nil {
		return
	}

	if sysAdminModel.ID <= 0 {
		err = errors.New("管理员不存在")
		return
	}

	if phone != "" {
		var phoneExistCount int64
		err = db.GetDb().Model(sysAdminModel).Where("phone = ?", phone).Where("id != ?", id).Count(&phoneExistCount).Error
		if err != nil {
			return
		}

		if phoneExistCount > 0 {
			err = errors.New("该号码已被绑定")
			return
		}
		sysAdminModel.Phone = phone
	}

	if email != "" {
		var emailExistCount int64
		err = db.GetDb().Model(sysAdminModel).Where("email = ?", email).Where("id != ?", id).Count(&emailExistCount).Error
		if err != nil {
			return
		}

		if emailExistCount > 0 {
			err = errors.New("该邮箱已被绑定")
			return
		}

		sysAdminModel.Email = email
	}

	if realname != "" {
		sysAdminModel.Realname = realname
	}

	if email != "" {
		sysAdminModel.Email = email
	}

	if password != "" {
		sysAdminModel.Password = utils.BcryptHash(password)
	}

	// 事务更新数据
	_ = db.GetDb().Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Save(&sysAdminModel).Error
		if err != nil {
			return
		}
		return
	})

	return
}

func (s *sSysAdmin) SetStatus(input sys_request.AdminSetStatusReq) (err error) {
	var (
		id     = input.Id
		status = input.Status

		sysAdminModel model.SysAdmin
	)

	err = db.GetDb().First(&sysAdminModel, id).Error
	if err != nil {
		return
	}
	if int(status) == sysAdminModel.Status {
		return
	}
	sysAdminModel.Operator = utils.GetUserName(s.ctx)
	err = db.GetDb().Save(&sysAdminModel).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysAdmin) List(input sys_request.AdminListReq) (out sys_request.AdminListRes, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
	)

	model1 := s.GetQuery(input)
	err = model1.Count(&out.Total).Error
	if err != nil {
		return
	}
	err = model1.Offset((page - 1) * pageSize).Limit(pageSize).Scan(&out.List).Error
	if err != nil {
		return
	}
	for key, item := range out.List {
		var deptInfo model.SysDept
		db.GetDb().Where("id = ?", item.DeptID).Find(&deptInfo, item.DeptID)
		out.List[key].DeptInfo = deptInfo
	}
	return
}

func (s *sSysAdmin) GetQuery(input sys_request.AdminListReq) *gorm.DB {
	var (
		username = input.Username
		realname = input.Realname
		email    = input.Email
		phone    = input.Phone

		sysAdminModel model.SysAdmin
	)

	model := db.GetDb().Model(sysAdminModel)
	if username != "" {
		model.Where("username like '%?%'", username)
	}
	if realname != "" {
		model.Where("username like '%?%'", realname)
	}

	if email != "" {
		model.Where("email = ?", email)
	}

	if phone != "" {
		model.Where("phone = ?", phone)
	}

	return model
}

func (s *sSysAdmin) GetAdminInfo() (adminModel model.SysAdmin, err error) {
	adminModel = utils.GetAdminInfo(s.ctx)
	if adminModel.ID <= 0 {
		token := utils.GetAuthorization(s.ctx)
		if token == "" {
			err = errors.New("参数异常")
			return
		}
		var mc *my_jwt.MyClaims
		mc, err = my_jwt.NewJWT().ParseToken(token)
		if err != nil {
			return
		}
		adminModel.ID = uint(mc.BaseClaims.Id)
	}
	err = db.GetDb().First(&adminModel, adminModel.ID).Error
	if err != nil {
		return
	}
	if adminModel.Status != 1 {
		err = errors.New("账号状态异常")
		return
	}
	//fmt.Printf("adminModel:  %+v\n", adminModel)
	return
}
