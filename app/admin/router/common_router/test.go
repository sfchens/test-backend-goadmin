package common_router

import (
	"csf/app/admin/apis/common_apis"
	"github.com/gin-gonic/gin"
)

func registerTestRouter(r *gin.RouterGroup) {
	api := common_apis.NewTestApi()
	r1 := r.Group("/test")
	{
		r1.POST("/index", api.Index)
		r1.POST("/index2", api.Index2)
	}
}
