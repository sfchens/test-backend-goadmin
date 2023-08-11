package live_service

import (
	"csf/app/admin/request/live_request"
	"csf/common/mysql/model"
	"csf/library/easy_db"
	"csf/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sBackdropService struct {
	ctx *gin.Context
}

func NewBackdropService(ctx *gin.Context) *sBackdropService {
	return &sBackdropService{ctx: ctx}
}

var liveType = map[int]string{
	1: "直播背景",
}

func (s *sBackdropService) AddOrEdit(input live_request.BackdropAddOrEditReq) (err error) {
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
	liveBackdropModel.Operator = utils.GetUserName(s.ctx)

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

func (s *sBackdropService) List(input live_request.BackdropListReq) (out live_request.BackdropListRes, err error) {
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

func (s *sBackdropService) GetQuery(input live_request.BackdropListReq) *gorm.DB {
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
