package live_apis

import (
	"csf/app/admin/request/live_request"
	"csf/app/admin/service/live_service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cBackdropApi struct{}

func NewBackdropApi() *cBackdropApi {
	return &cBackdropApi{}
}

// Add  添加背景
// @Summary 添加背景
// @Description 添加背景
// @Tags 直播背景
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/live/backdrop/add [post]
func (c *cBackdropApi) Add(ctx *gin.Context) {
	var (
		err error
		req live_request.BackdropAddOrEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = live_service.NewBackdropService(ctx).AddOrEdit(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// Edit  编辑背景
// @Summary 编辑背景
// @Description 编辑背景
// @Tags 直播背景
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/live/backdrop/edit [post]
func (c *cBackdropApi) Edit(ctx *gin.Context) {
	var (
		err error
		req live_request.BackdropAddOrEditReq
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	err = live_service.NewBackdropService(ctx).AddOrEdit(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// List  背景列表
// @Summary 背景列表
// @Description 背景列表
// @Tags 直播背景
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=live_request.BackdropListRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/live/backdrop/list [get]
func (c *cBackdropApi) List(ctx *gin.Context) {
	var (
		err error
		req live_request.BackdropListReq
		res live_request.BackdropListRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = live_service.NewBackdropService(ctx).List(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, res)
}
