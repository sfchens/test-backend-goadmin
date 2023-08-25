package user_service

import (
	"csf/app/admin/request/user_req"
	"github.com/gin-gonic/gin"
)

type sUser struct {
	ctx *gin.Context
}

func NewUserService(ctx *gin.Context) *sUser {
	return &sUser{ctx: ctx}
}

func (s *sUser) Add(input user_req.UserAddOrEditReq) (err error) {

	return
}
