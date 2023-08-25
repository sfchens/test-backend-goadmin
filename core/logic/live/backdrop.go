package live

import (
	"csf/core/mysql/model"
	"csf/core/query/live_query"
	"csf/core/service"
	"csf/library/easy_db"
	"csf/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	service.RegisterBackdrop(NewBackdropService())
}

type sBackdropService struct{}

func NewBackdropService() *sBackdropService {
	return &sBackdropService{}
}

var liveType = map[int]string{
	1: "直播背景",
}

func (s *sBackdropService) AddOrEdit(ctx *gin.Context, input live_query.BackdropAddOrEditInput) (err error) {
	var (
		id     = input.Id
		name   = input.Name
		types  = input.Type
		status = input.Status
		url    = input.Url

		liveBackdropModel model.LiveBackdrop
	)

	if id > 0 {
		err = easy_db.GetDb().Find(&liveBackdropModel, id).Error
		if err != nil {
			return
		}
	}

	if liveType[types] == "" {
		err = errors.New("直播类型参数异常")
		return
	}
	liveBackdropModel.Name = name
	liveBackdropModel.Type = types
	liveBackdropModel.Status = status
	liveBackdropModel.URL = url
	liveBackdropModel.Operator = utils.GetUserName(ctx)

	if id > 0 {
		err = easy_db.GetDb().Save(&liveBackdropModel).Error
	} else {
		err = easy_db.GetDb().Create(&liveBackdropModel).Error
	}
	if err != nil {
		return
	}
	return
}

func (s *sBackdropService) List(ctx *gin.Context, input live_query.BackdropListInput) (out live_query.BackdropListOut, err error) {
	var (
		page     = input.Page
		pageSize = input.PageSize
	)

	model := s.getQuery(input)
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

func (s *sBackdropService) getQuery(input live_query.BackdropListInput) *gorm.DB {
	var (
		name   = input.Name
		status = input.Status
		types  = input.Type
	)

	model := easy_db.GetDb().Model(model.LiveBackdrop{})

	if name != "" {
		model.Where("name like '%?%'", name)
	}

	if status != "" {
		model.Where("status =? ", status)
	}

	if types != "" {
		model.Where("type = ?", types)
	}

	return model
}
