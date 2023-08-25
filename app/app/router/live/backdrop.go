package live

import (
	"csf/app/admin/apis/live"
	"github.com/gin-gonic/gin"
)

func registerBackdropRouter(r *gin.RouterGroup) {
	api := live.NewBackdropApi()
	r1 := r.Group("/live/backdrop")
	{
		r1.POST("/add", api.Add)
		r1.POST("/edit", api.Edit)
		r1.GET("/list", api.List)
	}
}
