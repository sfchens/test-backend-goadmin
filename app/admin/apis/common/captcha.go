package common

import (
	"csf/app/admin/request/common_req"
	"csf/core/service"
	"csf/library/response"
	"github.com/gin-gonic/gin"
)

type captchaApi struct{}

func NewCaptchaApi() *captchaApi {
	return &captchaApi{}
}

// GetCaptcha 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Tags 验证码管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=common_req.GetOneRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/vi/sys/common/get_captcha [get]
func (c *captchaApi) GetCaptcha(ctx *gin.Context) {

	id, b64, err := service.NewCommonServiceGroup().CaptchaService.CreateCaptcha(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}

	res := common_req.GetOneRes{
		Id:   id,
		Path: b64,
	}
	response.SuccessWithData(ctx, res)
}
