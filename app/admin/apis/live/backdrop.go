package live

import (
	"csf/app/admin/request/live_req"
	"csf/core/query/live_query"
	"csf/core/service"
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
		req live_req.BackdropAddOrEditReq

		input live_query.BackdropAddOrEditInput
	)

	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = service.NewLiveServiceGroup().BackdropService.AddOrEdit(ctx, input)
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
		req live_req.BackdropAddOrEditReq

		input live_query.BackdropAddOrEditInput
	)

	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = service.NewLiveServiceGroup().BackdropService.AddOrEdit(ctx, input)
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
// @Success 200 {object} response.Response{data=live.BackdropListRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/live/backdrop/list [get]
func (c *cBackdropApi) List(ctx *gin.Context) {
	var (
		err   error
		req   live_req.BackdropListReq
		res   live_query.BackdropListOut
		input live_query.BackdropListInput
	)

	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	res, err = service.NewLiveServiceGroup().BackdropService.List(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, res)
}
