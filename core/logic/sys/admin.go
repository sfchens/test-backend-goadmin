package sys

import (
	"csf/core/mysql/model"
	"csf/core/query/sys_query"
	"csf/core/service"
	"csf/library/easy_db"
	"csf/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

func init() {
	service.RegisterNewAdmin(NewSysAdminService())
}

type sSysAdmin struct{}

func NewSysAdminService() *sSysAdmin {
	return &sSysAdmin{}
}

func (s *sSysAdmin) Add(ctx *gin.Context, input sys_query.AdminAddOrEditInput) (err error) {
	var (
		id            = input.Id
		sysAdminModel model.SysAdmin
	)
	sysAdminModel, err = s.DealAddOrEdit(ctx, input)
	if err != nil {
		return
	}

	if id > 0 {
		// 事务更新数据
		err = easy_db.GetDb().Transaction(func(tx *gorm.DB) (err error) {
			err = tx.Save(&sysAdminModel).Error
			if err != nil {
				return
			}
			return
		})
	} else {
		// 事务更新数据
		err = easy_db.GetDb().Transaction(func(tx *gorm.DB) (err error) {
			err = tx.Create(&sysAdminModel).Error
			if err != nil {
				return
			}
			return
		})
	}

	if err != nil {
		return
	}

	return
}

func (s *sSysAdmin) DealAddOrEdit(ctx *gin.Context, input sys_query.AdminAddOrEditInput) (sysAdminModel model.SysAdmin, err error) {
	var (
		id       = input.Id
		username = input.Username
		realname = input.Realname
		password = input.Password
		phone    = input.Phone
		email    = input.Email
		remark   = input.Remark
		sex      = input.Sex
		deptId   = input.DeptId
		status   = input.Status
		roleIds  = input.RoleIds
	)

	if id > 0 {
		err = easy_db.GetDb().Model(sysAdminModel).Find(&sysAdminModel, id).Error
		if err != nil {
			return
		}
	}

	var counts int64
	exitsModel := easy_db.GetDb().Model(sysAdminModel).Where("username=?", username)
	if id > 0 {
		exitsModel.Where("id != ?", id)
	}
	err = exitsModel.Count(&counts).Error
	if counts > 0 {
		err = errors.New("该账号已存在")
		return
	}

	if phone != "" {
		phoneExistModel := easy_db.GetDb().Model(sysAdminModel).Where("phone = ?", phone)
		if id > 0 {
			phoneExistModel.Where("id != ?", id)
		}

		var phoneExistCount int64
		err = phoneExistModel.Count(&phoneExistCount).Error
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
		emailExistModel := easy_db.GetDb().Model(sysAdminModel).Where("email = ?", email)
		if id > 0 {
			emailExistModel.Where("id != ?", id)
		}

		var emailExistCount int64
		err = emailExistModel.Count(&emailExistCount).Error
		if err != nil {
			return
		}

		if emailExistCount > 0 {
			err = errors.New("该邮箱已被绑定")
			return
		}
		sysAdminModel.Email = email
	}

	if password != "" {
		sysAdminModel.Password = utils.BcryptHash(password)
	}

	roleIdsNew := utils.IntToStringArray(roleIds)
	// admin表
	sysAdminModel.Remark = remark
	sysAdminModel.Sex = sex
	sysAdminModel.DeptID = deptId
	sysAdminModel.Status = status
	sysAdminModel.Username = username
	sysAdminModel.Realname = realname
	sysAdminModel.RoleIds = strings.Join(roleIdsNew, ",")
	return
}

func (s *sSysAdmin) SetStatus(ctx *gin.Context, input sys_query.AdminSetStatusInput) (err error) {
	var (
		id     = input.Id
		status = input.Status

		sysAdminModel model.SysAdmin
	)
	err = easy_db.GetDb().First(&sysAdminModel, id).Error
	if err != nil {
		return
	}
	if int(status) == sysAdminModel.Status {
		return
	}
	sysAdminModel.Status = int(status)
	sysAdminModel.Operator = utils.GetUserName(ctx)
	err = easy_db.GetDb().Save(&sysAdminModel).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysAdmin) List(ctx *gin.Context, input sys_query.AdminListInput) (out sys_query.AdminListOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize

		sysAdminList []model.SysAdmin
	)

	model1 := s.getQuery(ctx, input)
	err = model1.Count(&out.Total).Error
	if err != nil {
		return
	}
	err = model1.Offset((page - 1) * pageSize).Limit(pageSize).Scan(&sysAdminList).Error
	if err != nil {
		return
	}
	for _, item := range sysAdminList {
		var deptInfo model.SysDept
		easy_db.GetDb().Where("id = ?", item.DeptID).Find(&deptInfo, item.DeptID)

		var tmp sys_query.AdminListItem
		utils.StructToStruct(item, &tmp)
		tmp.DeptInfo = deptInfo
		tmp.RoleIds = utils.StringToIntArray(strings.Split(item.RoleIds, ","))

		// 权限
		var roleList []model.SysRole
		easy_db.GetDb().Model(model.SysRole{}).Where("id in (?)", tmp.RoleIds).Scan(&roleList)

		var textRole []string
		for _, val := range roleList {
			textRole = append(textRole, fmt.Sprintf(" %v ", val.Name))
		}
		tmp.RoleIdsText = strings.Join(textRole, ",")
		out.List = append(out.List, tmp)
	}
	return
}

func (s *sSysAdmin) getQuery(ctx *gin.Context, input sys_query.AdminListInput) *gorm.DB {
	var (
		username = input.Username
		realname = input.Realname
		email    = input.Email
		phone    = input.Phone
		status   = input.Status

		sysAdminModel model.SysAdmin
	)

	model := easy_db.GetDb().Model(sysAdminModel)
	if username != "" {
		model.Where(fmt.Sprintf("username like '%%%v%%'", username))
	}
	if realname != "" {
		model.Where(fmt.Sprintf("realname like '%%%v%%'", realname))
	}

	if email != "" {
		model.Where("email = ?", email)
	}

	if phone != "" {
		model.Where("phone = ?", phone)
	}

	if status != -1 {
		model.Where("status = ?", status)
	}
	return model
}

func (s *sSysAdmin) GetAdminInfo(ctx *gin.Context) (adminModel model.SysAdmin, err error) {
	var userInfo utils.UserInfo
	userInfo, err = utils.GetUserInfo(ctx)
	if err != nil {
		return
	}
	err = easy_db.GetDb().First(&adminModel, userInfo.Id).Error
	if err != nil {
		return
	}
	if adminModel.Status != 1 {
		err = errors.New("账号状态异常")
		return
	}
	return
}

func (s *sSysAdmin) ResetPwd(ctx *gin.Context, input sys_query.AdminResetPwdInput) (err error) {
	var (
		id = input.Id

		sysAdmin model.SysAdmin
	)

	err = easy_db.GetDb().Model(sysAdmin).Find(&sysAdmin, id).Error
	if err != nil {
		return
	}

	sysAdmin.Password = utils.BcryptHash("123456")
	err = easy_db.GetDb().Save(&sysAdmin).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysAdmin) DeleteBatch(ctx *gin.Context, input sys_query.AdminDeleteBatchInput) (err error) {
	var (
		ids = input.Ids

		sysAdminList []model.SysAdmin
	)

	tx := easy_db.GetDb().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Model(model.SysAdmin{}).Where("id in (?)", strings.Join(utils.IntToStringArray(ids), ",")).Scan(&sysAdminList).Error
	if err != nil {
		return
	}
	for _, item := range sysAdminList {
		err = easy_db.GetDb().Where("id =?", item.ID).Delete(&model.SysAdmin{}).Error
		if err != nil {
			break
		}
	}
	if err != nil {
		return
	}
	return
}

func (s *sSysAdmin) SetRole(ctx *gin.Context, input sys_query.AdminSetRoleInput) (err error) {

	var (
		id      = input.Id
		roleIds = input.RoleIds

		sysAdminModel model.SysAdmin
	)

	err = easy_db.GetDb().Model(sysAdminModel).Find(&sysAdminModel, id).Error
	if err != nil {
		return
	}

	sysAdminModel.RoleIds = strings.Join(utils.IntToStringArray(roleIds), ",")
	err = easy_db.GetDb().Save(&sysAdminModel).Error
	if err != nil {
		return
	}
	return
}
