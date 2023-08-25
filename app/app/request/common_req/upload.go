package common_req

import "mime/multipart"

type UploadPictureReq struct {
	Filename string                `json:"filename" form:"filename" description:"文件名称"`
	File     *multipart.FileHeader `json:"file" form:"file" validate:"required" msg:"required:请上传图片" description:"上传文件"`
	Path     string                `json:"path" form:"path" description:"保存路径"`
}

type UploadPictureRes struct {
	Md5Str   string `json:"md5_str" form:"md5_str" description:"md5"`
	Filename string `json:"filename" form:"filename" description:"名称"`
	Path     string `json:"path" form:"path" description:"保存路径"`
	Type     string `json:"type" form:"type" description:"图片类型"`
}

type UploadPictureMultiReq struct {
	Filename []string                `json:"filename" form:"filename" description:"文件名称"`
	File     []*multipart.FileHeader `json:"file" form:"file" validate:"required" msg:"required:请上传文件" description:"上传文件"`
	Path     []string                `json:"path" form:"path" description:"保存路径"`
}

type UploadPictureMultiRes struct {
	Md5Str   string `json:"md5_str" form:"md5_str" description:"md5"`
	Filename string `json:"filename" form:"filename" description:"名称"`
	Path     string `json:"path" form:"path" description:"保存路径"`
	Status   int    `json:"status" form:"status" description:"上传结果"`
	ErrMsg   string `json:"err_msg" form:"err_msg" description:"错误信息"`
}

type UploadAddPictureReq struct {
	Filename string                `json:"filename" form:"filename" description:"文件名称"`
	File     *multipart.FileHeader `json:"file" form:"file" validate:"required" msg:"required:请上传文件" description:"上传文件"`
	Path     string                `json:"path" form:"path" description:"保存路径"`
}

type UploadAddPictureRes struct {
	Md5Str   string `json:"md5_str" form:"md5_str" description:"md5"`
	Filename string `json:"filename" form:"filename" description:"名称"`
	Path     string `json:"path" form:"path" description:"保存路径"`
}

type UploadEditPictureReq struct {
	File *multipart.FileHeader `json:"file" form:"file"`
}

type UploadEditPictureRes struct {
}

type UploadVideoReq struct {
	Filename string                `json:"filename" form:"filename" description:"文件名称"`
	File     *multipart.FileHeader `json:"file" form:"file" validate:"required" msg:"required:请上传视频" description:"上传文件"`
	Path     string                `json:"path" form:"path" description:"保存路径"`
}

type UploadVideoRes struct {
	Md5Str   string `json:"md5_str" form:"md5_str" description:"md5"`
	Filename string `json:"filename" form:"filename" description:"名称"`
	Path     string `json:"path" form:"path" description:"保存路径"`
	Type     string `json:"type" form:"type" description:"视频类型"`
}
