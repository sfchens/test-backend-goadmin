package live

import (
	"csf/app/admin/request/live_req"
	"csf/core/query/live_query"
	"csf/core/service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cVideoApi struct{}

func NewVideoApi() *cVideoApi {
	return &cVideoApi{}
}

// Add  新增视频
// @Summary 新增视频
// @Description 新增视频
// @Tags 直播视频管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/live/video/add [post]
func (c *cVideoApi) Add(ctx *gin.Context) {
	var (
		err error
		req live_req.VideoAddOrEditReq

		input live_query.VideoAddOrEditInput
	)
	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = service.NewLiveServiceGroup().VideoService.AddOrEdit(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// Edit  编辑视频
// @Summary 编辑视频
// @Description 编辑视频
// @Tags 直播视频管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/live/video/edit [post]
func (c *cVideoApi) Edit(ctx *gin.Context) {
	var (
		err error
		req live_req.VideoAddOrEditReq

		input live_query.VideoAddOrEditInput
	)
	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	err = service.NewLiveServiceGroup().VideoService.AddOrEdit(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.Success(ctx)
}

// List  视频列表
// @Summary 视频列表
// @Description 视频列表
// @Tags 视频管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=live.VideoListRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/live/video/list [get]
func (c *cVideoApi) List(ctx *gin.Context) {
	var (
		err error
		req live_req.VideoListReq

		input live_query.VideoListInput
		res   live_query.VideoListOut
	)
	err = utils.BindParams(ctx, &req, &input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = service.NewLiveServiceGroup().VideoService.List(ctx, input)

	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, res)
}
