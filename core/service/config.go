package service

import (
	"csf/core/query/config_query"
	"github.com/gin-gonic/gin"
)

var localConfigService configServiceGroup

func NewConfigServiceGroup() configServiceGroup {
	return localConfigService
}

type configServiceGroup struct {
	ConfigService iConfig
	SwitchService iSwitch
}

type (
	iConfig interface {
		List(ctx *gin.Context, input config_query.ConfigListInput) (out config_query.ConfigListOut, err error)
		Add(ctx *gin.Context, input config_query.ConfigAddInput) (err error)
		Edit(ctx *gin.Context, input config_query.ConfigEditInput) (err error)
		GetOne(ctx *gin.Context, input config_query.ConfigGetOneInput) (out config_query.SysConfig, err error)
		DealJson(key string, dataJson string) (data interface{}, err error)
		Delete(ctx *gin.Context, input config_query.ConfigDeleteInput) (err error)
		SetStatus(ctx *gin.Context, input config_query.ConfigSetStatusInput) (err error)
	}

	iSwitch interface {
		AddOrEdit(ctx *gin.Context, input config_query.SwitchAddOrEditInput) (err error)
		List(ctx *gin.Context, input config_query.SwitchListInput) (out config_query.SwitchListOut, err error)
		Delete(ids []int) (err error)
		SetStatus(input config_query.SwitchSetStatusInput) (err error)
	}
)

func RegisterConfig(i iConfig) {
	localConfigService.ConfigService = i
}

func RegisterSwitch(i iSwitch) {
	localConfigService.SwitchService = i
}
