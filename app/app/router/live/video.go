package live

import (
	"csf/app/admin/apis/live"
	"github.com/gin-gonic/gin"
)

func registerVideoRouter(r *gin.RouterGroup) {
	api := live.NewVideoApi()
	r1 := r.Group("/live/video")
	{
		r1.POST("/add", api.Add)
		r1.POST("/edit", api.Edit)
		r1.GET("/list", api.List)
	}
}
