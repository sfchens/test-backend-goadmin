package common

import (
	"csf/core/mysql/model"
	"csf/core/query/common_query"
	"csf/core/service"
	"csf/library/easy_db"
	"csf/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
)

func init() {
	service.RegisterNewUpload(NewUploadService())
}
func NewUploadService() *sUploadService {
	return &sUploadService{}
}

type sUploadService struct {
	ctx *gin.Context
}

func (s *sUploadService) AddPicture(ctx *gin.Context, input common_query.UploadPictureInput) (out common_query.UploadPictureOut, err error) {
	var (
		uploadPicture model.ComPicture
	)
	out, err = s.UploadPicture(ctx, input)
	if err != nil {
		return
	}

	utils.StructToStruct(out, &uploadPicture)
	uploadPicture.Type = 1
	err = easy_db.GetDb().Create(uploadPicture).Error
	if err != nil {
		return
	}
	return
}

func (s *sUploadService) UploadPicture(ctx *gin.Context, input common_query.UploadPictureInput) (out common_query.UploadPictureOut, err error) {
	var (
		file     = input.File
		filename = input.Filename
		paths    = input.Path
	)
	extName := path.Ext(file.Filename) //获取后缀名
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
		".md":   true,
	}

	picType := extName[1:]

	if !allowExtMap[extName] {
		err = errors.New(fmt.Sprintf("暂不支持上传 %s 结尾的图片", extName))
		return
	}

	if filename == "" {
		filename = file.Filename
	}

	nameInfo := utils.GetFileName(filename)

	fmt.Printf("filename: %v； %+v\n", filename, nameInfo)
	filenameMd5 := utils.Md5(nameInfo.Name)
	var dir string
	dir, err = utils.GetPicTureDir(paths)
	if err != nil {
		return
	}

	newFileName := fmt.Sprintf("%v.%v", filenameMd5, nameInfo.ExtName)
	saveDir := path.Join(dir, newFileName)
	err = s.ctx.SaveUploadedFile(file, saveDir)
	if err != nil {
		return
	}

	out = common_query.UploadPictureOut{
		Md5Str:   filenameMd5,
		Filename: filename,
		Url:      utils.GetBaseUrl(saveDir),
		Type:     picType,
	}
	return
}

func (s *sUploadService) UploadPictureMulti(ctx *gin.Context, input common_query.UploadPictureMultiInput) (out []common_query.UploadPictureMultiOut) {
	var (
		fileNames = input.Filename
		files     = input.File
		paths     = input.Path
	)
	for i, file := range files {
		var (
			filename string
			path     string

			res = common_query.UploadPictureMultiOut{}
		)
		filename = file.Filename
		if len(fileNames) > 0 && fileNames[i] != "" {
			filename = fileNames[i]
		}

		if len(fileNames) > 0 && paths[i] != "" {
			path = paths[i]
		}

		var in = common_query.UploadPictureInput{
			Filename: filename,
			Path:     path,
			File:     file,
		}
		out1, err1 := s.UploadPicture(ctx, in)
		utils.StructToStruct(out1, &res)

		if err1 != nil {
			res.Filename = filename
			res.Path = path
			res.ErrMsg = err1.Error()
			res.Status = 2
		} else {
			res.Status = 1
			res.Filename = out1.Filename
			res.Path = out1.Url
			res.Md5Str = out1.Md5Str
		}
		out = append(out, res)
	}
	return
}

func (s *sUploadService) UploadVideo(ctx *gin.Context, input common_query.UploadVideoInput) (out common_query.UploadVideoOut, err error) {
	var (
		file     = input.File
		filename = input.Filename
		paths    = input.Path
	)
	extName := path.Ext(file.Filename) //获取后缀名
	allowExtMap := map[string]bool{
		".mp4": true,
	}
	if !allowExtMap[extName] {
		err = errors.New(fmt.Sprintf("暂不支持上传 %s 结尾的视频", extName))
		return
	}
	if filename == "" {
		filename = file.Filename
	}

	var dir string
	dir, err = utils.GetVideoDir(paths)
	if err != nil {
		return
	}

	saveDir := path.Join(dir, filename)
	err = s.ctx.SaveUploadedFile(file, saveDir)
	if err != nil {
		return
	}

	out = common_query.UploadVideoOut{
		Md5Str:   utils.Md5(filename),
		Filename: filename,
		Path:     saveDir,
		Type:     extName[1:],
	}
	return
}

func (s *sUploadService) EditPicture(ctx *gin.Context, input common_query.UploadEditPictureInput) (out common_query.UploadEditPictureOut, err error) {

	return
}
