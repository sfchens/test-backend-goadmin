package service

import (
	"csf/core/query/common_query"
	"github.com/gin-gonic/gin"
)

var localCommonService commonServiceGroup

func NewCommonServiceGroup() commonServiceGroup {
	return localCommonService
}

type commonServiceGroup struct {
	CaptchaService iCaptcha
	UploadService  iUpload
}

type (
	iCaptcha interface {
		CreateCaptcha(ctx *gin.Context) (id, b64s string, err error)
		Verify(ctx *gin.Context, id, code string, clear bool) bool
	}

	iUpload interface {
		AddPicture(ctx *gin.Context, input common_query.UploadPictureInput) (out common_query.UploadPictureOut, err error)
		UploadPicture(ctx *gin.Context, input common_query.UploadPictureInput) (out common_query.UploadPictureOut, err error)
		UploadPictureMulti(ctx *gin.Context, input common_query.UploadPictureMultiInput) (out []common_query.UploadPictureMultiOut)
		UploadVideo(ctx *gin.Context, input common_query.UploadVideoInput) (out common_query.UploadVideoOut, err error)
		EditPicture(ctx *gin.Context, input common_query.UploadEditPictureInput) (out common_query.UploadEditPictureOut, err error)
	}
)

func RegisterNewCaptcha(i iCaptcha) {
	localCommonService.CaptchaService = i
}

func RegisterNewUpload(i iUpload) {
	localCommonService.UploadService = i
}
