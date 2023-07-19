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
)

type sSwitchService struct {
	ctx *gin.Context
}

func NewSwitchService(ctx *gin.Context) *sSwitchService {
	return &sSwitchService{ctx: ctx}
}

func (s *sSwitchService) AddOrEdit(input sys_request.SwitchAddOrEditReq) (err error) {
	var (
		id      = input.Id
		name    = input.Name
		typeKey = input.TypeKey
		status  = input.Status
		remark  = input.Remark

		existCount    int64
		sysWitchModel model.SysSwitch
	)

	existModel := db.GetDb().Model(model.SysSwitch{})
	if id > 0 {
		existModel.Where("id != ?", id)
	}
	err = existModel.Where("type_key = ?", typeKey).Count(&existCount).Error
	if err != nil {
		return
	}

	if existCount > 0 {
		err = errors.New(fmt.Sprintf("键名%s已存在", typeKey))
	}
	sysWitchModel.TypeKey = typeKey
	sysWitchModel.Name = name
	sysWitchModel.Status = int(status)
	sysWitchModel.Remark = remark
	sysWitchModel.Operator = utils.GetUserName(s.ctx)

	if id > 0 {
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
	)
	model := s.GetQuery(input)

	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model.Offset((page - 1) * pageSize).Limit(pageSize).Scan(&out.List).Error
	if err != nil {
		return
	}
	return
}

func (s *sSwitchService) GetQuery(input sys_request.SwitchListReq) *gorm.DB {
	var (
		name    = input.Name
		typeKey = input.TypeKey
	)

	model := db.GetDb().Model(model.SysSwitch{})

	if name != "" {
		model.Where("name like '%?%'", name)
	}
	if typeKey != "" {
		model.Where("type_key =?", typeKey)
	}
	return model
}
