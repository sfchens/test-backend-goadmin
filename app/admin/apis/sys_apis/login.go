package sys_apis

import (
	"csf/app/admin/model/sys_model"
	"csf/app/admin/request/sys_request"
	"csf/app/admin/service/common_service"
	"csf/app/admin/service/sys_service"
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
// @Param raw body     sys_request.LoginReq true "请求参数"
// @Success 200 {object} response.Response{data=sys_request.LoginRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/login [post]
func (c cSysLoginApi) Login(ctx *gin.Context) {
	var (
		err error
		req sys_request.LoginReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	if easy_config.Viper.Get("app.mode") != "dev" && !common_service.NewComCaptchaService(ctx).Verify(req.CaptChaId, req.Captcha, true) {
		response.FailWithMessage(ctx, "验证码验证失败")
		return
	}

	var (
		inputLogin = sys_model.LoginInput{
			Username: req.Username,
			Password: req.Password,
		}

		loginRes sys_request.LoginRes
	)
	loginRes, err = sys_service.NewSysLoginService(ctx).Login(inputLogin)
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
// @Param raw body     sys_request.LoginReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/login_info [get]
func (c cSysLoginApi) LoginInfo(ctx *gin.Context) {
	var (
		err error

		req sys_request.ConfigGetOneReq
		res sys_request.ConfigGetOneRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	var out sys_model.SysConfig
	out, err = sys_service.NewSysConfigService(ctx).GetOne(req)
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
// @Param raw body     sys_request.LoginReq true "请求参数"
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/sys/logout [post]
func (c cSysLoginApi) Logout(ctx *gin.Context) {
	var (
		err error
	)
	err = sys_service.NewSysLoginService(ctx).Logout()
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
	}
	response.SuccessWithMessage(ctx, "退出成功")
}
