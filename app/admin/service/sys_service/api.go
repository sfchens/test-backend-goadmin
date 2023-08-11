package sys_service

import (
	"bytes"
	"csf/app/admin/request/sys_request"
	"csf/common/mysql/model"
	"csf/library/easy_db"
	"csf/library/global"
	"csf/utils"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io/ioutil"
	"strings"
)

type sSysApiService struct {
	ctx *gin.Context
}

func NewSysApiService(ctx *gin.Context) *sSysApiService {
	return &sSysApiService{ctx: ctx}
}

func (s *sSysApiService) List(input sys_request.ApiListReq) (out sys_request.ApiListRes, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
	)

	model1 := s.GetQuery(input)

	err = model1.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = model1.Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Scan(&out.List).Error
	if err != nil {
		return
	}

	return
}

func (s *sSysApiService) GetQuery(input sys_request.ApiListReq) *gorm.DB {
	var (
		tag    = input.Tag
		title  = input.Title
		path   = input.Path
		method = input.Method
	)

	model1 := easy_db.GetDb().Model(model.SysApi{})

	if tag != "" {
		model1.Where(fmt.Sprintf("tags like '%%%v%%'", tag))
	}

	if title != "" {
		model1.Where(fmt.Sprintf("title like '%%%v%%'", title))
	}

	if path != "" {
		model1.Where(fmt.Sprintf("path like '%%%v%%'", path))
	}
	if method != "" {
		model1.Where("method = ?", strings.ToUpper(method))
	}
	return model1
}

// Refresh 刷新接口
func (s *sSysApiService) Refresh() (err error) {
	// 获取所有路由信息
	routers := global.GinEngine.Routes()
	// 可在此处增加配置路径前缀的if判断，只对代码生成的自建应用进行定向的接口名称填充
	jsonFile, _ := ioutil.ReadFile("docs/swagger.json")
	jsonData, _ := simplejson.NewFromReader(bytes.NewReader(jsonFile))
	for _, route := range routers {
		var (
			sysApiModel model.SysApi
		)
		json := jsonData.Get("paths").Get(route.Path).Get(strings.ToLower(route.Method))
		title, _ := json.Get("summary").String()
		tags, _ := json.Get("tags").StringArray()
		if strings.Contains(route.Path, "/swagger/") || title == "" {
			continue
		}

		sysApiModel.Path = route.Path
		sysApiModel.Method = route.Method
		sysApiModel.Handle = route.Handler
		sysApiModel.Title = title
		sysApiModel.Tags = strings.Join(tags, ",")
		sysApiModel.Operator = global.OperatorSystem

		err = easy_db.GetDb().Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&sysApiModel).Error
		if err != nil {
			//break
		}
	}
	return
}

func (s *sSysApiService) AddOrEdit(input sys_request.ApiEditReq) (err error) {
	var (
		id     = input.Id
		tags   = input.Tags
		title  = input.Title
		method = input.Method

		sysApiModel model.SysApi
	)

	if id > 0 {
		err = easy_db.GetDb().Find(&sysApiModel, id).Error
		if err != nil {
			return
		}
	}

	sysApiModel.Tags = tags
	sysApiModel.Title = title
	sysApiModel.Method = method
	sysApiModel.Operator = utils.GetUserName(s.ctx)

	if sysApiModel.ID > 0 {
		err = easy_db.GetDb().Save(&sysApiModel).Error
	} else {
		sysApiModel.Path = input.Path
		sysApiModel.Handle = input.Handle
		err = easy_db.GetDb().Create(&sysApiModel).Error
	}
	if err != nil {
		return
	}
	return
}

func (s *sSysApiService) GetTag(input sys_request.ApiGetTagReq) (out sys_request.ApiGetTagRes, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
		tag      = input.Tag
	)

	model := easy_db.GetDb().Model(model.SysApi{}).Group("tags").Select("tags")
	if tag != "" {
		model.Where("tags like %?%", tag)
	}
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

func (s *sSysApiService) DeleteMulti(input sys_request.ApiDeleteMultiReq) (err error) {

	var (
		ids = input.Ids

		idsStr          []string
		sysApiModel     model.SysApi
		sysApiModelList []model.SysApi
	)
	for _, v := range ids {
		idsStr = append(idsStr, fmt.Sprintf("%v", v))
	}
	err = easy_db.GetDb().Model(sysApiModel).
		Where(fmt.Sprintf("id in (%v)", strings.Join(idsStr, ","))).
		Scan(&sysApiModelList).Error
	if err != nil {
		return
	}

	tx := easy_db.GetDb().Begin()
	for _, v := range sysApiModelList {
		err = tx.Delete(&sysApiModel, v.ID).Error
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
