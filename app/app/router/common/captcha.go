package common

import (
	"csf/app/admin/apis/common"
	"github.com/gin-gonic/gin"
)

func registerCommonRouter(r *gin.RouterGroup) {

	captchaObj := common.NewCaptchaApi()
	r1 := r.Group("/captcha")
	{
		r1.GET("/get_one", captchaObj.GetCaptcha)
	}
}
