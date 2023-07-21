package sys_service

import (
	"csf/app/admin/request/sys_request"
	"csf/common/mysql/model"
	"csf/library/db"
	"csf/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type sSwitchService struct {
	ctx *gin.Context
}

func NewSwitchService(ctx *gin.Context) *sSwitchService {
	return &sSwitchService{ctx: ctx}
}

func (s *sSwitchService) AddOrEdit(input sys_request.SwitchAddOrEditReq) (err error) {
	var (
		id     = input.Id
		name   = input.Name
		key    = input.Key
		status = input.Status
		remark = input.Remark

		existCount    int64
		sysWitchModel model.SysSwitch
	)

	existModel := db.GetDb().Model(model.SysSwitch{})
	if id > 0 {
		existModel.Where("id != ?", id)
	}
	err = existModel.Where("`key` = ?", key).Count(&existCount).Error
	if err != nil {
		return
	}

	if existCount > 0 {
		err = errors.New(fmt.Sprintf("键名%s已存在", key))
	}
	if id > 0 {
		err = db.GetDb().Model(model.SysSwitch{}).Scan(&sysWitchModel).Error
		if err != nil {
			return
		}
	}
	sysWitchModel.Key = key
	sysWitchModel.Name = name
	sysWitchModel.Status = int(status)
	sysWitchModel.Remark = remark
	sysWitchModel.Operator = utils.GetUserName(s.ctx)

	if sysWitchModel.ID > 0 {
		err = db.GetDb().Save(&sysWitchModel).Error
	} else {
		err = db.GetDb().Create(&sysWitchModel).Error
	}

	if err != nil {
		return
	}

	return
}

func (s *sSwitchService) List(input sys_request.SwitchListReq) (out sys_request.SwitchListRes, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
		order    = input.Order
	)
	model := s.GetQuery(input)

	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model.Offset((page - 1) * pageSize).Limit(pageSize).Order(order).Scan(&out.List).Error
	if err != nil {
		return
	}
	return
}

func (s *sSwitchService) GetQuery(input sys_request.SwitchListReq) *gorm.DB {
	var (
		name = input.Name
		key  = input.Key
	)

	model := db.GetDb().Model(model.SysSwitch{})

	if name != "" {
		model.Where("name like '%?%'", name)
	}
	if key != "" {
		model.Where("`key` =?", key)
	}
	return model
}

func (s *sSwitchService) Delete(ids []int) (err error) {

	var (
		idsStr             []string
		sysSwitchModel     model.SysSwitch
		sysSwitchModelList []model.SysSwitch
	)
	for _, v := range ids {
		idsStr = append(idsStr, fmt.Sprintf("%v", v))
	}
	err = db.GetDb().Model(sysSwitchModel).
		Where(fmt.Sprintf("id in (%v)", strings.Join(idsStr, ","))).
		Scan(&sysSwitchModelList).Error
	if err != nil {
		return
	}

	tx := db.GetDb().Begin()
	for _, v := range sysSwitchModelList {
		if v.Status == 1 {
			err = errors.New(fmt.Sprintf("配置ID： %v 正在使用", v.ID))
			break
		}

		err = db.GetDb().Delete(&sysSwitchModel, v.ID).Error
		if err != nil {
			break
		}
	}
	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return
}

func (s *sSwitchService) SetStatus(input sys_request.SwitchSetStatusReq) (err error) {
	var (
		id     = input.Id
		status = input.Status

		sysSwitchModel model.SysSwitch
	)

	err = db.GetDb().Find(&sysSwitchModel, id).Error
	if err != nil {
		return
	}

	if sysSwitchModel.Status == status {
		err = errors.New("状态异常，请刷新后重试")
		return
	}
	sysSwitchModel.Status = status
	err = db.GetDb().Save(&sysSwitchModel).Error
	if err != nil {
		return
	}

	return
}