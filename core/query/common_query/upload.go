package common_query

import "mime/multipart"

type UploadPictureInput struct {
	Filename string                `json:"filename" form:"filename" description:"文件名称"`
	File     *multipart.FileHeader `json:"file" form:"file" validate:"required" msg:"required:请上传图片" description:"上传文件"`
	Path     string                `json:"path" form:"path" description:"保存路径"`
}
type UploadPictureOut struct {
	Md5Str   string `json:"md5_str" form:"md5_str" description:"md5"`
	Filename string `json:"filename" form:"filename" description:"名称"`
	Path     string `json:"path" form:"path" description:"保存路径"`
	Type     string `json:"type" form:"type" description:"图片类型"`
}

type UploadAddPictureInput struct {
	Filename string                `json:"filename" form:"filename" description:"文件名称"`
	File     *multipart.FileHeader `json:"file" form:"file" validate:"required" msg:"required:请上传文件" description:"上传文件"`
	Path     string                `json:"path" form:"path" description:"保存路径"`
}
type UploadAddPictureOut struct {
	Md5Str   string `json:"md5_str" form:"md5_str" description:"md5"`
	Filename string `json:"filename" form:"filename" description:"名称"`
	Path     string `json:"path" form:"path" description:"保存路径"`
}

type UploadPictureMultiInput struct {
	Filename []string                `json:"filename" form:"filename" description:"文件名称"`
	File     []*multipart.FileHeader `json:"file" form:"file" validate:"required" msg:"required:请上传文件" description:"上传文件"`
	Path     []string                `json:"path" form:"path" description:"保存路径"`
}
type UploadPictureMultiOut struct {
	Md5Str   string `json:"md5_str" form:"md5_str" description:"md5"`
	Filename string `json:"filename" form:"filename" description:"名称"`
	Path     string `json:"path" form:"path" description:"保存路径"`
	Status   int    `json:"status" form:"status" description:"上传结果"`
	ErrMsg   string `json:"err_msg" form:"err_msg" description:"错误信息"`
}

type UploadVideoInput struct {
	Filename string                `json:"filename" form:"filename" description:"文件名称"`
	File     *multipart.FileHeader `json:"file" form:"file" validate:"required" msg:"required:请上传视频" description:"上传文件"`
	Path     string                `json:"path" form:"path" description:"保存路径"`
}

type UploadVideoOut struct {
	Md5Str   string `json:"md5_str" form:"md5_str" description:"md5"`
	Filename string `json:"filename" form:"filename" description:"名称"`
	Path     string `json:"path" form:"path" description:"保存路径"`
	Type     string `json:"type" form:"type" description:"视频类型"`
}

type UploadEditPictureInput struct {
	File *multipart.FileHeader `json:"file" form:"file"`
}
type UploadEditPictureOut struct {
}
