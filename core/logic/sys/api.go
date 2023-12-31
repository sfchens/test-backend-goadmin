package sys

import (
	"bytes"
	"csf/core/mysql/model"
	"csf/core/query/sys_query"
	"csf/core/service"
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

func init() {
	service.RegisterNewApi(NewSysApiService())
}

type sSysApiService struct{}

func NewSysApiService() *sSysApiService {
	return &sSysApiService{}
}

func (s *sSysApiService) List(ctx *gin.Context, input sys_query.ApiListInput) (out sys_query.ApiListOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
	)

	m := s.getQuery(ctx, input)

	err = m.Count(&out.Total).Error
	if err != nil {
		return
	}

	err = m.Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Scan(&out.List).Error
	if err != nil {
		return
	}
	return
}

func (s *sSysApiService) getQuery(ctx *gin.Context, input sys_query.ApiListInput) *gorm.DB {
	var (
		tag    = input.Tag
		title  = input.Title
		path   = input.Path
		method = input.Method
	)

	model1 := easy_db.GetDb().Model(model.SysAPI{})

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
			sysApiModel model.SysAPI
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

func (s *sSysApiService) AddOrEdit(ctx *gin.Context, input sys_query.ApiEditInput) (err error) {
	var (
		id     = input.Id
		tags   = input.Tags
		title  = input.Title
		method = input.Method

		sysApiModel model.SysAPI
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
	sysApiModel.Operator = utils.GetUserName(ctx)

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

func (s *sSysApiService) GetTag(ctx *gin.Context, input sys_query.ApiGetTagInput) (out sys_query.ApiGetTagOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
		tag      = input.Tag
	)

	model := easy_db.GetDb().Model(model.SysAPI{}).Group("tags").Select("tags")
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

func (s *sSysApiService) DeleteMulti(ctx *gin.Context, input sys_query.ApiDeleteMultiInput) (err error) {

	var (
		ids = input.Ids

		idsStr          []string
		sysApiModel     model.SysAPI
		sysApiModelList []model.SysAPI
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
