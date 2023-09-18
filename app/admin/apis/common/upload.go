package common

import (
	"csf/app/admin/request/common_req"
	"csf/core/query/common_query"
	"csf/core/service"
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
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  common_req.UploadAddPictureReq true "请求参数"
// @Success 200 {object} response.Response{data=common_req.UploadAddPictureRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/upload/add_picture [post]
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
	var (
		input common_query.UploadPictureInput
		out   common_query.UploadPictureOut
	)
	utils.StructToStruct(req, &input)
	out, err = service.NewCommonServiceGroup().UploadService.AddPicture(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(out, &res)
	response.SuccessWithData(ctx, res)
}

// EditPicture 编辑图片
// @Summary 编辑图片
// @Description 编辑图片
// @Tags 上传管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  common_req.UploadEditPictureReq true "请求参数"
// @Success 200 {object} response.Response{data=common_query.UploadEditPictureOut} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/upload/edit_picture [post]
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
	var (
		input common_query.UploadEditPictureInput
		out   common_query.UploadEditPictureOut
	)
	utils.StructToStruct(req, &input)
	out, err = service.NewCommonServiceGroup().UploadService.EditPicture(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(out, &res)
	response.SuccessWithData(ctx, res)
}

// UploadPicture 上传图片
// @Summary 上传图片
// @Description 上传图片
// @Tags 上传管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  common_req.UploadPictureReq true "请求参数"
// @Success 200 {object} response.Response{data=common_req.UploadPictureRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/upload/picture [post]
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

	var (
		input common_query.UploadPictureInput
		out   common_query.UploadPictureOut
	)
	utils.StructToStruct(req, &input)
	out, err = service.NewCommonServiceGroup().UploadService.UploadPicture(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	utils.StructToStruct(out, &res)
	response.SuccessWithData(ctx, res)
}

// UploadPictureMulti 批量上传图片
// @Summary 批量上传图片
// @Description 批量上传图片
// @Tags 上传管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  common_req.UploadPictureMultiReq true "请求参数"
// @Success 200 {object} response.Response{data=[]common_req.UploadPictureMultiRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/upload/picture_multi [post]
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
	var (
		input common_query.UploadPictureMultiInput
		out   []common_query.UploadPictureMultiOut
	)
	utils.StructToStruct(req, &input)
	out = service.NewCommonServiceGroup().UploadService.UploadPictureMulti(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	utils.StructToStruct(out, &res)
	response.SuccessWithData(ctx, res)
}

// UploadVideo 上传视频
// @Summary 上传视频
// @Description 上传视频
// @Tags 上传管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body  common_req.UploadVideoReq true "请求参数"
// @Success 200 {object} response.Response{data=common_req.UploadVideoRes} "code错误码 msg操作信息 data返回信息"
// @Router /admin/v1/upload/video [post]
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
	var (
		input common_query.UploadVideoInput
		out   common_query.UploadVideoOut
	)
	utils.StructToStruct(req, &input)
	out, err = service.NewCommonServiceGroup().UploadService.UploadVideo(ctx, input)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	utils.StructToStruct(out, &res)
	response.SuccessWithData(ctx, res)
}
