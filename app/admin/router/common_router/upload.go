package common_router

import (
	"csf/app/admin/apis/common_apis"
	"github.com/gin-gonic/gin"
)

func registerUploadRouter(r *gin.RouterGroup) {
	api := common_apis.NewUploadApi()
	r1 := r.Group("/upload")
	{
		r1.POST("/add_picture", api.AddPicture)
		r1.POST("/edit_picture", api.EditPicture)
		r1.POST("/picture", api.UploadPicture)
		r1.POST("/picture_multi", api.UploadPictureMulti)
		r1.POST("/video", api.UploadVideo)
	}
}
