package common_router

import (
	"csf/app/admin/apis/common_apis"
	"github.com/gin-gonic/gin"
)

func registerCommonRouter(r *gin.RouterGroup) {

	captchaObj := common_apis.NewCaptchaApi()
	r1 := r.Group("/captcha")
	{
		r1.GET("/get_one", captchaObj.GetCaptcha)
	}
}
