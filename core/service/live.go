package service

import (
	"csf/core/query/live_query"
	"github.com/gin-gonic/gin"
)

var localLiveService liveServiceGroup

func NewLiveServiceGroup() liveServiceGroup {
	return localLiveService
}

type liveServiceGroup struct {
	BackdropService iBackdrop
	VideoService    iVideo
}

type (
	iBackdrop interface {
		AddOrEdit(ctx *gin.Context, input live_query.BackdropAddOrEditInput) (err error)
		List(ctx *gin.Context, input live_query.BackdropListInput) (out live_query.BackdropListOut, err error)
	}

	iVideo interface {
		AddOrEdit(ctx *gin.Context, input live_query.VideoAddOrEditInput) (err error)
		List(ctx *gin.Context, input live_query.VideoListInput) (out live_query.VideoListOut, err error)
	}
)

func RegisterBackdrop(i iBackdrop) {
	localLiveService.BackdropService = i
}

func RegisterVideo(i iVideo) {
	localLiveService.VideoService = i
}
