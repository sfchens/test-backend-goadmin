package common

import (
	"csf/app/admin/request/common_req"
	"csf/app/admin/service/common_service"
	"csf/library/response"
	"csf/utils"
	"github.com/gin-gonic/gin"
)

type cUploadApi struct{}

func NewUploadApi() *cUploadApi {
	return &cUploadApi{}
}

// AddPicture 新增图片
// @Summary 新增图片
// @Description 新增图片
// @Tags 上传管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=common.UploadAddPictureRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/upload/add_picture [post]
func (c *cUploadApi) AddPicture(ctx *gin.Context) {
	var (
		err error
		req common_req.UploadAddPictureReq
		res common_req.UploadAddPictureRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = common_service.NewUploadService(ctx).AddPicture(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// EditPicture 编辑图片
// @Summary 编辑图片
// @Description 编辑图片
// @Tags 上传管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=common.UploadEditPictureRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/upload/edit_picture [post]
func (c *cUploadApi) EditPicture(ctx *gin.Context) {
	var (
		err error
		req common_req.UploadEditPictureReq
		res common_req.UploadEditPictureRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	res, err = common_service.NewUploadService(ctx).EditPicture(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}

// UploadPicture 上传图片
// @Summary 上传图片
// @Description 上传图片
// @Tags 上传管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=common.UploadPictureRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/upload/picture [post]
func (c *cUploadApi) UploadPicture(ctx *gin.Context) {
	var (
		err error

		req common_req.UploadPictureReq
		res common_req.UploadPictureRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res, err = common_service.NewUploadService(ctx).UploadPicture(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, res)
}

// UploadPictureMulti 批量上传图片
// @Summary 批量上传图片
// @Description 批量上传图片
// @Tags 上传管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]common.UploadPictureMultiRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/upload/picture_multi [post]
func (c *cUploadApi) UploadPictureMulti(ctx *gin.Context) {
	var (
		err error

		req common_req.UploadPictureMultiReq
		res []common_req.UploadPictureMultiRes
	)

	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	res = common_service.NewUploadService(ctx).UploadPictureMulti(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	response.SuccessWithData(ctx, res)
}

// UploadVideo 上传视频
// @Summary 上传视频
// @Description 上传视频
// @Tags 上传管理
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=common.UploadVideoRes} "code错误码 msg操作信息 data返回信息"
// @Router /api/v1/upload/video [post]
func (c *cUploadApi) UploadVideo(ctx *gin.Context) {
	var (
		err error
		req common_req.UploadVideoReq
		res common_req.UploadVideoRes
	)
	err = utils.BindParams(ctx, &req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	res, err = common_service.NewUploadService(ctx).UploadVideo(req)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	response.SuccessWithData(ctx, res)
}
