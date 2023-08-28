package live

import (
	"csf/core/mysql/model"
	"csf/core/query/live_query"
	"csf/core/service"
	"csf/library/easy_db"
	"csf/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	service.RegisterVideo(NewVideoService())
}

type sVideoService struct {
}

func NewVideoService() *sVideoService {
	return &sVideoService{}
}

func (s *sVideoService) AddOrEdit(ctx *gin.Context, input live_query.VideoAddOrEditInput) (err error) {
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
	liveVideo.Operator = utils.GetUserName(ctx)

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

func (s *sVideoService) List(ctx *gin.Context, input live_query.VideoListInput) (out live_query.VideoListOut, err error) {
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

func (s *sVideoService) getQuery(input live_query.VideoListInput) *gorm.DB {
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
