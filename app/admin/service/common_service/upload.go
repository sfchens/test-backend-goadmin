package common_service

import (
	"csf/app/admin/request/common_req"
	"csf/common/mysql/model"
	"csf/library/easy_db"
	"csf/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
)

type sUploadService struct {
	ctx *gin.Context
}

func NewUploadService(ctx *gin.Context) *sUploadService {
	return &sUploadService{ctx: ctx}
}

func (s *sUploadService) AddPicture(input common_req.UploadAddPictureReq) (out common_req.UploadAddPictureRes, err error) {

	var (
		uploadPictureReq common_req.UploadPictureReq
		uploadPictureRes common_req.UploadPictureRes

		uploadPicture model.ComPicture
	)
	uploadPictureRes, err = s.UploadPicture(uploadPictureReq)
	if err != nil {
		return
	}

	utils.StructToStruct(uploadPictureRes, &uploadPicture)
	uploadPicture.Type = 1
	err = easy_db.GetDb().Create(uploadPicture).Error
	if err != nil {
		return
	}
	return
}

func (s *sUploadService) EditPicture(input common_req.UploadEditPictureReq) (out common_req.UploadEditPictureRes, err error) {

	return
}

func (s *sUploadService) UploadPicture(input common_req.UploadPictureReq) (out common_req.UploadPictureRes, err error) {
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

	var dir string
	dir, err = utils.GetPicTureDir(paths)
	if err != nil {
		return
	}

	saveDir := path.Join(dir, filename)
	err = s.ctx.SaveUploadedFile(file, saveDir)
	if err != nil {
		return
	}

	out = common_req.UploadPictureRes{
		Md5Str:   utils.Md5(filename),
		Filename: filename,
		Path:     saveDir,
		Type:     picType,
	}
	return
}

func (s *sUploadService) UploadPictureMulti(input common_req.UploadPictureMultiReq) (out []common_req.UploadPictureMultiRes) {
	var (
		fileNames = input.Filename
		files     = input.File
		paths     = input.Path
	)
	for i, file := range files {
		var (
			filename string
			path     string

			res = common_req.UploadPictureMultiRes{}
		)
		filename = file.Filename
		if len(fileNames) > 0 && fileNames[i] != "" {
			filename = fileNames[i]
		}

		if len(fileNames) > 0 && paths[i] != "" {
			path = paths[i]
		}

		var in = common_req.UploadPictureReq{
			Filename: filename,
			Path:     path,
			File:     file,
		}
		out1, err1 := s.UploadPicture(in)
		utils.StructToStruct(out1, &res)

		if err1 != nil {
			res.Filename = filename
			res.Path = path
			res.ErrMsg = err1.Error()
			res.Status = 2
		} else {
			res.Status = 1
			res.Filename = out1.Filename
			res.Path = out1.Path
			res.Md5Str = out1.Md5Str
		}
		out = append(out, res)
	}
	return
}

func (s *sUploadService) UploadVideo(input common_req.UploadVideoReq) (out common_req.UploadVideoRes, err error) {
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

	out = common_req.UploadVideoRes{
		Md5Str:   utils.Md5(filename),
		Filename: filename,
		Path:     saveDir,
		Type:     extName[1:],
	}
	return
}
