package common

import (
	"csf/app/admin/apis/common"
	"github.com/gin-gonic/gin"
)

func registerTestRouter(r *gin.RouterGroup) {
	api := common.NewTestApi()
	r1 := r.Group("/test")
	{
		r1.POST("/index", api.Index)
		r1.POST("/index2", api.Index2)
	}
}
