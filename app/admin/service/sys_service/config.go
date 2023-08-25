package sys_service

import (
	"csf/app/admin/model/sys_model"
	"csf/app/admin/request/sys_req"
	"csf/core/mysql/model"
	"csf/library/easy_db"
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

func (s *sSysConfig) List(input sys_req.ConfigListReq) (out sys_req.ConfigListRes, err error) {
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
		var itemTmp sys_req.ConfigGetOneRes
		utils.StructToStruct(item, &itemTmp)
		var config interface{}
		config, _ = s.DealJson(item.Key, item.Config)
		itemTmp.Config = config
		out.List = append(out.List, itemTmp)
	}
	return
}

func (s *sSysConfig) GetQuery(input sys_req.ConfigListReq) *gorm.DB {
	var (
		name    = input.Name
		key     = input.Key
		keyName = input.KeyName

		sysConfigModel model.SysConfig
	)

	model := easy_db.GetDb().Model(sysConfigModel)

	if name != "" {
		model.Where(fmt.Sprintf("name like '%%%v%%'", name))
	}

	if key != "" {
		model.Where(fmt.Sprintf("`key` like '%%%v%%'", key))
	}
	if keyName != "" {
		model.Where(fmt.Sprintf("`config` like '%%%v%%'", keyName))
	}
	return model
}

func (s *sSysConfig) Add(input sys_req.ConfigAddReq) (err error) {
	var (
		name           = input.Name
		config         = input.Config
		key            = input.Key
		remark         = input.Remark
		isOpen         = input.IsOpen
		sysConfigModel model.SysConfig
	)
	if !json.Valid([]byte(config)) {
		err = errors.New("配置数据格式异常")
		return
	}

	var counts int64
	err = easy_db.GetDb().Model(sysConfigModel).Where("id=?", key).Count(&counts).Error
	if err != nil {
		return
	}
	if counts > 0 {
		err = errors.New("该类型已存在")
		return
	}
	sysConfigModel.Remark = remark
	sysConfigModel.IsOpen = uint(isOpen)
	sysConfigModel.Name = name
	sysConfigModel.Config = config
	sysConfigModel.Key = key
	sysConfigModel.Operator = utils.GetUserName(s.ctx)
	err = easy_db.GetDb().Create(&sysConfigModel).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysConfig) Edit(input sys_req.ConfigEditReq) (err error) {
	var (
		id     = input.Id
		name   = input.Name
		key    = input.Key
		config = input.Config
		isOpen = input.IsOpen
		remark = input.Remark

		sysConfigModel model.SysConfig
	)

	err = s.CheckConfig(key, config)
	if err != nil {
		return
	}

	err = easy_db.GetDb().First(&sysConfigModel, id).Error
	if err != nil {
		return
	}
	if sysConfigModel.ID <= 0 {
		err = errors.New("参数异常")
		return
	}
	configTmp, _ := json.Marshal(config)
	sysConfigModel.Key = key
	sysConfigModel.Name = name
	sysConfigModel.Config = string(configTmp)
	sysConfigModel.IsOpen = uint(isOpen)
	sysConfigModel.Remark = remark

	sysConfigModel.Operator = utils.GetUserName(s.ctx)
	_ = easy_db.GetDb().Save(&sysConfigModel).Error
	if err != nil {
		return
	}

	return
}

func (s *sSysConfig) CheckConfig(key string, config map[string]string) (err error) {
	switch key {
	case "BASE_CONFIG":
		if config["sys_app_logo"] == "" {
			err = errors.New("系统logo不为空")
		}
		if config["sys_app_name"] == "" {
			err = errors.New("网站名称不为空")
		}
	default:
		err = errors.New("参数异常")
	}
	return
}

func (s *sSysConfig) GetOne(input sys_req.ConfigGetOneReq) (out sys_model.SysConfig, err error) {
	var (
		sysConfigModel model.SysConfig
	)
	if input.Id <= 0 && input.Key == "" {
		err = errors.New("参数异常")
		return
	}
	model := easy_db.GetDb().Model(sysConfigModel)
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

func (s *sSysConfig) Delete(input sys_req.ConfigDeleteReq) (err error) {
	var (
		ids            = input.Ids
		idsStr         []string
		sysConfigList  []model.SysConfig
		sysConfigModel model.SysConfig
	)

	for _, v := range ids {
		idsStr = append(idsStr, fmt.Sprintf("%v", v))
	}

	err = easy_db.GetDb().Model(sysConfigModel).
		Where(fmt.Sprintf("id in (%v)", strings.Join(idsStr, ","))).
		Scan(&sysConfigList).Error
	if err != nil {
		return
	}

	tx := easy_db.GetDb().Begin()
	for _, v := range sysConfigList {
		if v.IsOpen == 1 {
			err = errors.New(fmt.Sprintf("配置ID： %v 正在使用", v.ID))
			break
		}

		err = easy_db.GetDb().Delete(&sysConfigModel, v.ID).Error
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

func (s *sSysConfig) SetStatus(input sys_req.ConfigSetStatusReq) (err error) {
	var (
		id     = input.Id
		isOpen = input.IsOpen

		sysConfigModel model.SysConfig
	)

	err = easy_db.GetDb().Find(&sysConfigModel, id).Error
	if err != nil {
		return
	}

	if sysConfigModel.IsOpen == uint(isOpen) {
		err = errors.New("状态异常，刷新后重试")
		return
	}
	sysConfigModel.IsOpen = uint(isOpen)
	sysConfigModel.Operator = utils.GetUserName(s.ctx)
	err = easy_db.GetDb().Save(&sysConfigModel).Error
	if err != nil {
		return
	}
	return
}
