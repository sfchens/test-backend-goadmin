package sys_service

import (
	"csf/app/admin/model/sys_model"
	"csf/app/admin/request/sys_request"
	"csf/common/mysql/model"
	"csf/library/db"
	"csf/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type sSysConfig struct {
	ctx *gin.Context
}

func NewSysConfigService(ctx *gin.Context) *sSysConfig {
	return &sSysConfig{ctx: ctx}
}

func (s *sSysConfig) List(input sys_request.ConfigListReq) (out sys_request.ConfigListRes, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
		order    = input.Order

		sysConfigList []model.SysConfig
	)

	model := s.GetQuery(input)
	err = model.Count(&out.Total).Error
	if err != nil {
		return
	}
	err = model.Offset((page - 1) * pageSize).Limit(pageSize).Order(order).Scan(&sysConfigList).Error
	if err != nil {
		return
	}
	for _, item := range sysConfigList {
		var itemTmp sys_request.ConfigGetOneRes
		utils.StructToStruct(item, &itemTmp)
		var config interface{}
		if item.Type == 2 {
			config, _ = s.DealJson(item.Key, item.Config)
		} else {
			config = item.Config
		}
		itemTmp.Config = config
		out.List = append(out.List, itemTmp)
	}
	return
}

func (s *sSysConfig) GetQuery(input sys_request.ConfigListReq) *gorm.DB {
	var (
		name  = input.Name
		types = input.Type1
		key   = input.Key

		sysConfigModel model.SysConfig
	)

	model := db.GetDb().Model(sysConfigModel)

	if name != "" {
		model.Where(fmt.Sprintf("name like '%%%v%%'", name))
	}

	if types != "" {
		model.Where("type = ?", types)
	}

	if key != "" {
		model.Where(fmt.Sprintf("`key` like '%%%v%%'", key))
	}
	return model
}

func (s *sSysConfig) Add(input sys_request.ConfigAddReq) (err error) {
	var (
		name           = input.Name
		config         = input.Config
		key            = input.Key
		types          = input.Type // 1value, 2需要转json格式
		remark         = input.Remark
		isOpen         = input.IsOpen
		sysConfigModel model.SysConfig
	)
	//if !json.Valid([]byte(config)) {
	//	err = errors.New("配置数据格式异常")
	//	return
	//}
	if types == 1 {

	}
	var counts int64
	err = db.GetDb().Model(sysConfigModel).Where("id=?", key).Count(&counts).Error
	if err != nil {
		return
	}
	if counts > 0 {
		err = errors.New("该类型已存在")
		return
	}
	sysConfigModel.Type = types
	sysConfigModel.Remark = remark
	sysConfigModel.IsOpen = uint(isOpen)
	sysConfigModel.Name = name
	sysConfigModel.Config = fmt.Sprintf("%v", config)
	sysConfigModel.Key = key
	sysConfigModel.Operator = utils.GetUserName(s.ctx)
	err = db.GetDb().Create(&sysConfigModel).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysConfig) Edit(input sys_request.ConfigEditReq) (err error) {
	var (
		id     = input.Id
		name   = input.Name
		config = input.Config
		type1  = input.Type
		isOpen = input.IsOpen
		remark = input.Remark

		sysConfigModel model.SysConfig
	)

	if !json.Valid([]byte(config)) {
		err = errors.New("配置数据格式异常")
		return
	}

	err = db.GetDb().First(&sysConfigModel, id).Error
	if err != nil {
		return
	}
	if sysConfigModel.ID <= 0 {
		err = errors.New("参数异常")
		return
	}
	sysConfigModel.Name = name
	sysConfigModel.Config = config
	sysConfigModel.Type = type1
	sysConfigModel.IsOpen = uint(isOpen)
	sysConfigModel.Remark = remark

	sysConfigModel.Operator = utils.GetUserName(s.ctx)
	_ = db.GetDb().Save(&sysConfigModel).Error
	if err != nil {
		return
	}

	return
}

func (s *sSysConfig) GetOne(input sys_request.ConfigGetOneReq) (out sys_model.SysConfig, err error) {
	var (
		sysConfigModel model.SysConfig
	)
	if input.Id <= 0 && input.Key == "" {
		err = errors.New("参数异常")
		return
	}
	model := db.GetDb().Model(sysConfigModel)
	if input.Id > 0 {
		model.Where("id = ?", input.Id)
	}
	if input.Key != "" {
		model.Where("`key` = ?", input.Key)
	}
	err = model.Scan(&sysConfigModel).Error
	if err != nil {
		return
	}
	utils.StructToStruct(sysConfigModel, &out)
	var configJson interface{}
	configJson, err = s.DealJson(sysConfigModel.Key, sysConfigModel.Config)
	if err != nil {
		return
	}
	out.Config = configJson
	return
}

func (s *sSysConfig) DealJson(key string, dataJson string) (data interface{}, err error) {
	switch key {
	case "BASE_CONFIG": // 基础配置
		err = json.Unmarshal([]byte(dataJson), &data)
	default:
		err = errors.New("类型异常")
		return
	}
	if err != nil {
		return
	}
	return
}

func (s *sSysConfig) Delete(input sys_request.ConfigDeleteReq) (err error) {
	var (
		ids            = input.Ids
		idsStr         []string
		sysConfigList  []model.SysConfig
		sysConfigModel model.SysConfig
	)

	for _, v := range ids {
		idsStr = append(idsStr, fmt.Sprintf("%v", v))
	}

	err = db.GetDb().Model(sysConfigModel).
		Where(fmt.Sprintf("id in (%v)", strings.Join(idsStr, ","))).
		Scan(&sysConfigList).Error
	if err != nil {
		return
	}

	tx := db.GetDb().Begin()
	for _, v := range sysConfigList {
		if v.IsOpen == 1 {
			err = errors.New(fmt.Sprintf("配置ID： %v 正在使用", v.ID))
			break
		}

		err = db.GetDb().Delete(&sysConfigModel, v.ID).Error
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

func (s *sSysConfig) SetStatus(input sys_request.ConfigSetStatusReq) (err error) {
	var (
		id     = input.Id
		isOpen = input.IsOpen

		sysConfigModel model.SysConfig
	)

	err = db.GetDb().Find(&sysConfigModel, id).Error
	if err != nil {
		return
	}

	if sysConfigModel.IsOpen == uint(isOpen) {
		err = errors.New("状态异常，刷新后重试")
		return
	}
	sysConfigModel.IsOpen = uint(isOpen)
	sysConfigModel.Operator = utils.GetUserName(s.ctx)
	err = db.GetDb().Save(&sysConfigModel).Error
	if err != nil {
		return
	}
	return
}
