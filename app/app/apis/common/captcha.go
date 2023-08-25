package common

import (
	"csf/app/admin/request/common_req"
	"csf/app/admin/service/common_service"
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
// @Success 200 {object} response.Response{data=common.GetOneRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/vi/sys/common/get_captcha [get]
func (c *captchaApi) GetCaptcha(ctx *gin.Context) {

	id, b64, err := common_service.NewComCaptchaService(ctx).CreateCaptcha()
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}

	res := common_req.GetOneRes{
		Id:   id,
		Path: b64,
	}
	response.SuccessWithData(ctx, res)
}
