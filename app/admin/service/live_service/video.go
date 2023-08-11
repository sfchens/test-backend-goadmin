package live_service

import (
	"csf/app/admin/request/live_request"
	"csf/common/mysql/model"
	"csf/library/easy_db"
	"csf/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sVideoService struct {
	ctx *gin.Context
}

func NewVideoService(ctx *gin.Context) *sVideoService {
	return &sVideoService{ctx: ctx}
}

func (s *sVideoService) AddOrEdit(input live_request.VideoAddOrEditReq) (err error) {
	var (
		id     = input.Id
		name   = input.Name
		types  = input.Type
		status = input.Status
		url    = input.Url

		liveVideo model.LiveVideo
	)
	if id > 0 {
		err = easy_db.GetDb().Find(&liveVideo, id).Error
		if err != nil {
			return
		}
	}
	liveVideo.Name = name
	liveVideo.Type = types
	liveVideo.Status = status
	liveVideo.URL = url
	liveVideo.Operator = utils.GetUserName(s.ctx)

	if id > 0 {
		err = easy_db.GetDb().Save(&liveVideo).Error
	} else {
		err = easy_db.GetDb().Create(&liveVideo).Error
	}
	if err != nil {
		return
	}
	return
}

func (s *sVideoService) List(input live_request.VideoListReq) (out live_request.VideoListRes, err error) {
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

func (s *sVideoService) GetQuery(input live_request.VideoListReq) *gorm.DB {
	var (
		name   = input.Name
		status = input.Status
		types  = input.Type
	)

	model := easy_db.GetDb().Model(model.LiveVideo{})

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
