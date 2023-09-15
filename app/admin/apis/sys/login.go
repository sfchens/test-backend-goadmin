package sys

import (
	"csf/app/admin/request/sys_req"
	"csf/core/query/config_query"
	"csf/core/query/login_query"
	"csf/core/service"
	"csf/library/easy_config"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cSysLoginApi struct{}

func NewSysLogin() *cSysLoginApi {
	return &cSysLoginApi{}
}

// Login  登录
// @Summary 登录
// @Description 登录
// @Tags 管理员登录管理
// @Accept application/json
// @Produce application/json
// @Param object body     sys_req.LoginReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_req.LoginRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/login [post]
func (c cSysLoginApi) Login(ctx *gin.Context) {
	var (
		err error
		req sys_req.LoginReq
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if easy_config.Config.App.Mode != "dev" && !service.NewCommonServiceGroup().CaptchaService.Verify(ctx, req.CaptChaId, req.Captcha, true) {
		response.FailWithMessage(ctx, "验证码验证失败")
		return
	}

	var (
		inputLogin = login_query.AdminLoginInput{
			Username: req.Username,
			Password: req.Password,
		}

		loginRes login_query.AdminLoginOut
	)
	loginRes, err = service.NewLoginServiceGroup().AdminLoginService.Login(ctx, inputLogin)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, loginRes)
}

// LoginInfo  登录信息
// @Summary 登录信息
// @Description 登录信息
// @Tags 管理员登录管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query     sys_req.ConfigGetOneReq true "请求参数"
// @Success 200 {object} response.Response{data=config_query.ConfigGetOneOut} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/login_info [get]
func (c cSysLoginApi) LoginInfo(ctx *gin.Context) {
	var (
		err error

		req sys_req.ConfigGetOneReq

		input config_query.ConfigGetOneInput
		res   config_query.ConfigGetOneOut
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(req, &input)
	var out config_query.SysConfig
	out, err = service.NewConfigServiceGroup().ConfigService.GetOne(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	utils.StructToStruct(out, &res)
	response.SuccessWithData(ctx, res)
}

// Logout  退出
// @Summary 退出
// @Description 退出
// @Tags 管理员登录管理
// @Accept application/json
// @Produce application/json
// @Param raw body     sys_req.LoginReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/sys/logout [post]
func (c cSysLoginApi) Logout(ctx *gin.Context) {
	var (
		err error
	)
	err = service.NewLoginServiceGroup().AdminLoginService.Logout(ctx)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	response.SuccessWithMessage(ctx, "退出成功")
}
