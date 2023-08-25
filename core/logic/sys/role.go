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
	"gorm.io/gorm/clause"
	"strings"
)

func init() {
	service.RegisterNewRole(NewSysRoleService())
}

type sSysRoleService struct{}

func NewSysRoleService() *sSysRoleService {
	return &sSysRoleService{}
}

func (s *sSysRoleService) AddOrEdit(ctx *gin.Context, input sys_query.RoleAddOrEditInput) (err error) {
	var (
		id      = input.Id
		name    = input.Name
		key     = input.Key
		status  = input.Status
		sort    = input.Sort
		remark  = input.Remark
		menuIds = input.MenuIds

		sysRoleModel model.SysRole
	)
	var exitCount int64
	exitsModel := easy_db.GetDb().Model(sysRoleModel)

	if id > 0 {

		exitsModel.Where("id != ?", id)

		err = easy_db.GetDb().Model(sysRoleModel).Find(&sysRoleModel, id).Error
		if err != nil {
			return
		}
	}

	err = exitsModel.Where("`key`=?", key).Count(&exitCount).Error
	if err != nil {
		return
	}

	if exitCount > 0 {
		err = errors.New(fmt.Sprintf("权限标识%v已存在", key))
		return
	}

	tx := easy_db.GetDb().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	sysRoleModel.Status = status
	sysRoleModel.Name = name
	sysRoleModel.Key = key
	sysRoleModel.Sort = int64(sort)
	sysRoleModel.Remark = remark
	sysRoleModel.MenuIds = strings.Join(utils.IntToStringArray(menuIds), ",")
	if id > 0 {
		err = tx.Save(&sysRoleModel).Error
	} else {
		err = tx.Create(&sysRoleModel).Error
	}
	if err != nil {
		return
	}
	if key != "superAdmin" {
		err = s.SaveRoleMenu(ctx, tx, int(sysRoleModel.ID), menuIds)
	}

	return
}

func (s *sSysRoleService) SaveRoleMenu(ctx *gin.Context, tx *gorm.DB, roleId int, menuIds []int) (err error) {

	err = tx.Where("role_id = ?", roleId).Delete(&model.SysRoleMenu{}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	for _, menuId := range menuIds {
		var sysRoleMenuModel model.SysRoleMenu
		sysRoleMenuModel.MenuID = int64(menuId)
		sysRoleMenuModel.RoleID = int64(roleId)
		err = tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&sysRoleMenuModel).Error
		if err != nil {
			break
		}
	}
	if err != nil {
		tx.Rollback()
		return
	}
	return
}

func (s *sSysRoleService) List(ctx *gin.Context, input sys_query.RoleListInput) (out sys_query.RoleListOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize

		sysRoleList []model.SysRole
	)

	model := s.getQuery(input)

	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model.Offset((page - 1) * pageSize).Limit(pageSize).Scan(&sysRoleList).Error
	if err != nil {
		return
	}

	for _, item := range sysRoleList {
		var roleItem sys_query.RoleListItem
		utils.StructToStruct(item, &roleItem)
		menuIds := strings.Split(item.MenuIds, ",")
		roleItem.MenuIds = utils.StringToIntArray(menuIds)
		out.List = append(out.List, roleItem)
	}

	return
}

func (s *sSysRoleService) getQuery(input sys_query.RoleListInput) *gorm.DB {
	var (
		name   = input.Name
		key    = input.Key
		status = input.Status

		sysRoleModel model.SysRole
	)
	model := easy_db.GetDb().Model(sysRoleModel)
	if name != "" {
		model.Where(fmt.Sprintf("name like '%%%v%%'", name))
	}

	if key != "" {
		model.Where(fmt.Sprintf("`key` like '%%%v%%'", key))
	}

	if status != 0 {
		model.Where("status = ?", status)
	}
	return model
}

func (s *sSysRoleService) DeleteBatch(ctx *gin.Context, input sys_query.RoleDeleteBatchInput) (err error) {

	var (
		ids = input.Ids

		errNew []string
	)
	for _, id := range ids {
		newInput := sys_query.RoleDeleteInput{
			Id: id,
		}
		err = s.Delete(ctx, newInput)
		if err != nil {
			errNew = append(errNew, fmt.Sprintf("序号： %v 删除失败, 错误信息： %v", id, err.Error()))
		}
	}

	if len(errNew) > 0 {
		err = errors.New(strings.Join(errNew, "\n,"))
		return
	}

	return
}

func (s *sSysRoleService) Delete(ctx *gin.Context, input sys_query.RoleDeleteInput) (err error) {
	var (
		id               = input.Id
		sysRoleModel     model.SysRole
		sysRoleMenuModel model.SysRoleMenu
	)

	err = easy_db.GetDb().First(&sysRoleModel, id).Error
	if err != nil {
		return
	}

	tx := easy_db.GetDb().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	err = tx.Where("id =?", id).Delete(&sysRoleModel).Error
	if err != nil {
		return
	}

	err = tx.Where("role_id =?", id).Delete(&sysRoleMenuModel).Error
	if err != nil {
		return
	}
	return
}
