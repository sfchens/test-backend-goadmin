package live_router

import (
	"csf/app/admin/apis/live_apis"
	"github.com/gin-gonic/gin"
)

func registerVideoRouter(r *gin.RouterGroup) {
	api := live_apis.NewVideoApi()
	r1 := r.Group("/live/video")
	{
		r1.POST("/add", api.Add)
		r1.POST("/edit", api.Edit)
		r1.GET("/list", api.List)
	}
}
